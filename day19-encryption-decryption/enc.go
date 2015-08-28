package cryp

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Common actions are printed to stdout.  These can be silenced
// by setting LogOutput = ioutil.Discard
var LogOutput io.Writer = os.Stdout

// Encrypt takes data and a key and outputs encrypted data and any possible errors
// The key can be any length or empty (not recommended).  The data can be any length
// or empty.  A SHA-512/256 key is generated from the supplied key ensuring the
// 32 byte AES-256 key length requirement is met.  The data is compressed using
// gzip prior to encryption.  Raw byte output will need to be hex/base64 encoded
// before it is printable.
func Encrypt(data []byte, key []byte) ([]byte, error) {

	// generate a 32 byte key from the variable length key supplied
	// Sum512_256 == 32 byte key == aes256
	aes256Key := sha512.Sum512_256(key)

	block, err := aes.NewCipher(aes256Key[:])
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	// we always encrypt the key with the data
	// to ensure that enough data is encrypted.
	// this allows encrypting empty or small strings safely.
	// it is trimmed from the output in decrypt
	if _, err := w.Write(aes256Key[:]); err != nil {
		return nil, err
	}
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+buf.Len())

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], buf.Bytes())

	return ciphertext, nil

}

// EncryptDirFiles takes a directory and a key and searches, recursively,
// for any files to encrypt.  Each file's file name, size, mode, mod time, and contents
// are written in tar format and passed to Encrypt. A new file is created
// that is named the SHA-256 checksum of the encrypted output.  This effectively hides
// the file name and provides a checksum to ensure contents are intact. All directories,
// symlinks, named pipes, sockets, and devices are left as-is.
func EncryptDirFiles(dir string, key []byte) error {

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		// Skip these types
		// ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
		if info.Mode()&os.ModeType != 0 {
			return nil
		}

		start := time.Now()
		fmt.Fprint(LogOutput, "Encrypting ", path)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Create a new tar archive.
		buf := &bytes.Buffer{}
		tw := tar.NewWriter(buf)

		hdr := &tar.Header{
			Name:    info.Name(),        // string    // name of header file entry
			Mode:    int64(info.Mode()), // int64     // permission and mode bits
			Size:    info.Size(),        // int64     // length in bytes
			ModTime: info.ModTime(),     // time.Time // modified time
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write(data); err != nil {
			return err
		}
		if err := tw.Close(); err != nil {
			return err
		}

		// encrypt tar archive
		encrypted_data, err := Encrypt(buf.Bytes(), key)
		if err != nil {
			return err
		}

		// create file name from hash of encrypted data
		h := sha256.New()
		h.Write(encrypted_data)
		new_file_name := hex.EncodeToString(h.Sum(nil))

		// create new file
		new_file_path := filepath.Join(filepath.Dir(path), new_file_name)
		new_file, err := os.OpenFile(new_file_path, os.O_WRONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC, 0400)
		if err != nil {
			return err
		}
		if n, err := new_file.Write(encrypted_data); err != nil {
			return err
		} else if n != len(encrypted_data) {
			return io.ErrShortWrite
		}
		if err := new_file.Close(); err != nil {
			return err
		}

		// remove original file
		if err := os.Remove(path); err != nil {
			return err
		}

		fmt.Fprintln(LogOutput, " ...", time.Since(start))

		return nil

	})

}
