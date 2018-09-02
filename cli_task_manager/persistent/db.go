package persistent

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

const (
	dbPath = "resources/my.db"

	//bolt db bucket
	bucket = "task"
)

type Task struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Time      time.Time `json:"time"`
	Completed bool      `json:"completed"`
}

//don't forget to close connection
//defer db.Close()
func openDBConn() *bolt.DB {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AddTask(name string) bool {
	db := openDBConn()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			log.Fatalf("Create bucket operation failed\n %v", err)
		}
		id, err := buck.NextSequence()
		if err != nil {
			log.Fatalf("Failed creating auto incremented value for task\n %v", err)
		}
		t := Task{id, name, time.Now(), false}
		encoded, err := json.Marshal(t)
		if err != nil {
			log.Fatalf("Failed marshall task\n %v", err)
		}
		err = buck.Put(itob(int(id)), []byte(encoded))
		if err != nil {
			log.Fatalf("Failed store task to db\n %v", err)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed store task to db\n %v", err)
	}

	return true
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func ViewTasks() []Task {
	db := openDBConn()
	defer db.Close()
	tasks := make([]Task, 0)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			log.Fatalf("Failed get tasks from db\n Bucket is non-existed.")
		}
		c := b.Cursor()

		for _, v := c.First(); v != nil; _, v = c.Next() {
			t := Task{}
			err := json.Unmarshal([]byte(v), &t)
			if err != nil {
				log.Fatalf("Error unmarshalling task %v\n", err)
			}
			tasks = append(tasks, t)
		}

		return nil

	})
	if err != nil {
		log.Fatalf("Failed get tasks from db\n %v", err)
	}

	return tasks
}
