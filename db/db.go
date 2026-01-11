package db

import (
	"encoding/binary"
	"errors"

	bolt "go.etcd.io/bbolt"
)

var (
	tasksBucket = []byte("tasks")
)

func Open() (*bolt.DB, error) {
	return bolt.Open("tasks.db", 0600, nil)
}

func AddTask(db *bolt.DB, task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return err
		}

		id, _ := b.NextSequence()
		key := itob(id)

		return b.Put(key, []byte(task))
	})
}

func ListTasks(db *bolt.DB) (map[int]string, error) {
	tasks := make(map[int]string)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		if b == nil {
			return nil
		}

		i := 1
		return b.ForEach(func(k, v []byte) error {
			tasks[i] = string(v)
			i++
			return nil
		})
	})

	return tasks, err
}

func DeleteTask(db *bolt.DB, index int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		if b == nil {
			return errors.New("no tasks")
		}

		var targetKey []byte
		i := 1

		err := b.ForEach(func(k, _ []byte) error {
			if i == index {
				targetKey = k
			}
			i++
			return nil
		})

		if err != nil {
			return err
		}

		if targetKey == nil {
			return errors.New("task not found")
		}

		return b.Delete(targetKey)
	})
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
