package cryp

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var sha256HexRegexp = regexp.MustCompile(`^[a-f0-9]{64}$`)

// Decrypt takes data and a key and outputs decrypted data and any possible errors
// The key can be any length or empty (not recommended).  A SHA-512/256 key is generated
// from the supplied key ensuring the 32 byte AES-256 key length requirement is met.
// Once decrypted, the data is decompressed using gzip.
func Decrypt(data []byte, key []byte) ([]byte, error) {

	// should never happen
	if len(data) < aes.BlockSize {
		return nil, errors.New("insufficient data to decrypt")
	}

	// generate a 32 byte key from the variable length key supplied
	// Sum512_256 == 32 byte key == aes256
	aes256Key := sha512.Sum512_256(key)

	block, err := aes.NewCipher(aes256Key[:])
	if err != nil {
		return nil, err
	}

	iv := data[:aes.BlockSize]
	dectext := data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(dectext, dectext)

	r, err := gzip.NewReader(bytes.NewReader(dectext))
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// trim off key that was prepended during encryption
	return data[len(aes256Key):], nil

}

// DecryptDirFiles takes a directory and a key, and searches recursively,
// for any files that are named a SHA-256 checksum to decrypt.  It passes each
// file to Decrypt and parses the tar payload format. It attempts to restore
// the file to its original form (name, contents, mode, mod time) and returns any errors that
// occur.  Any files that do not match the SHA-256 checksum are left as-is.
func DecryptDirFiles(dir string, key []byte) error {

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		// continue into dirs
		if info.IsDir() {
			return nil
		}

		// expecting file name to be sha256 hex encoded hash of
		// 32ebb1abcc1c601ceb9c4e3c4faba0caa5b85bb98c4f1e6612c40faa528a91c9 (64 chars long)
		if !sha256HexRegexp.MatchString(filepath.Base(path)) {
			fmt.Fprintln(LogOutput, "Skipping", path)
			return nil
		}

		start := time.Now()
		fmt.Fprint(LogOutput, "Decrypting ", path)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		h := sha256.New()
		h.Write(data)
		data_hash := hex.EncodeToString(h.Sum(nil))
		if data_hash != filepath.Base(path) {
			return fmt.Errorf("Corruption detected in %s", path)
		}

		dec_data, err := Decrypt(data, key)
		if err != nil {
			return err
		}

		tr := tar.NewReader(bytes.NewReader(dec_data))
		hdr, err := tr.Next()
		if err != nil && err != io.EOF {
			return err
		}

		orig_file_path := filepath.Join(filepath.Dir(path), hdr.Name)
		orig_file, err := os.OpenFile(orig_file_path, os.O_WRONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC, os.FileMode(hdr.Mode))
		if err != nil {
			return err
		}
		if n, err := io.Copy(orig_file, tr); err != nil {
			orig_file.Close()
			return err
		} else if n != hdr.Size {
			orig_file.Close()
			return io.ErrShortWrite
		}
		if err := orig_file.Close(); err != nil {
			return err
		}
		if err := os.Chtimes(orig_file_path, time.Now(), hdr.ModTime); err != nil {
			return err
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		fmt.Fprintln(LogOutput, " ...", time.Since(start))

		return nil

	})

}
