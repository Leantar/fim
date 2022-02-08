package models

import (
	"encoding/hex"
	"fmt"
	"github.com/zeebo/blake3"
	"golang.org/x/sys/unix"
	"io"
	"os"
)

const (
	S_IFMT  = 0o0170000
	S_IFREG = 0o0100000
)

type FsObject struct {
	Path     string
	Hash     string
	Created  int64
	Modified int64
	Uid      uint32
	Gid      uint32
	Mode     uint32
}

func NewFsObject(path string) (FsObject, error) {
	var stat unix.Stat_t

	err := unix.Lstat(path, &stat)
	if err != nil {
		return FsObject{}, fmt.Errorf("failed to stat path: %w", err)
	}

	created, _ := stat.Ctim.Unix()
	modified, _ := stat.Mtim.Unix()

	obj := FsObject{
		Path:     path,
		Created:  created,
		Modified: modified,
		Uid:      stat.Uid,
		Gid:      stat.Gid,
		Mode:     stat.Mode,
	}

	// Check if file is regular
	if stat.Mode&S_IFMT == S_IFREG {
		obj.Hash, err = hashFile(path)
		if err != nil {
			return FsObject{}, err
		}
	}

	return obj, nil
}

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hasher := blake3.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
