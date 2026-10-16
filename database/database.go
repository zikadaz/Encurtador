package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zikadaz/encurtador/model"
	bolt "go.etcd.io/bbolt"
)

var DB *bolt.DB

func Connection() {
	// tenta criar /db na raiz do projeto
    if err := os.MkdirAll("../db", 0755); err != nil {
        log.Fatal(err)
    }
    var err error
	// abre uma conexao com o banco
    DB, err = bolt.Open(
        "../db/local.db",
        0600,
        &bolt.Options{Timeout: 1 * time.Second},
    )
    if err != nil {
        log.Fatal(err)
    }
	// cria um novo bucket
    err = DB.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("links"))
        return err
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("--> [Sys] : Banco de dados conectado.")
}

func Close() {
    if DB != nil {
        if err := DB.Close(); err != nil {
            log.Println(err)
        }
    }
}
// incrementa os clicks 
func IncrementClicks(code string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("links"))
		if bucket == nil {
			return fmt.Errorf("bucket 'links' não existe")
		}

		data := bucket.Get([]byte(code))
		if data == nil {
			return fmt.Errorf("código '%s' não encontrado", code)
		}

		var link model.Url
		if err := json.Unmarshal(data, &link); err != nil {
			return err
		}

		link.Clicks++

		updated, err := json.Marshal(link)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(code), updated)
	})
}

//  cria um novo registro no banco de dados
func AddInfoToDatabase(c /*code*/ string, s /*url*/ string) error {
	var clicks int  = 0
	model.Link = model.Url {
		Url: s,
		Code: c,
		Clicks: clicks,
	}
	data, err := json.Marshal(model.Link)
	if err != nil {
		return err
	}
	return DB.Update(func(tx *bolt.Tx) error{
		bucket := tx.Bucket([]byte("links"))
		if bucket == nil {
			return fmt.Errorf("Bucket 'links' nao existe.")
		}
		return bucket.Put([]byte(c), data)
	})
}
// consulta um objeto no banco de dados
func ReadInfoFromDatabase(code string) (map[string]model.Url, error) {
	result := make(map[string]model.Url)

	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("links"))
		if bucket == nil {
			return fmt.Errorf("bucket 'links' não existe")
		}

		// Procurar um código específico
		if code != "" {
			data := bucket.Get([]byte(code))
			if data == nil {
				return fmt.Errorf("código '%s' não encontrado", code)
			}

			var link model.Url
			if err := json.Unmarshal(data, &link); err != nil {
				return err
			}

			result[code] = link
			return nil
		}

		// Código vazio: listar todos
		return bucket.ForEach(func(k, v []byte) error {
			var link model.Url
			if err := json.Unmarshal(v, &link); err != nil {
				return err
			}

			result[string(k)] = link
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
func ClearDatabase() error {
	return DB.Update(func(tx *bolt.Tx) error {
		if err := tx.DeleteBucket([]byte("links")); err != nil {
			return err
		}

		_, err := tx.CreateBucket([]byte("links"))
		return err
	})
}