package cryp

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func init() {
	// discard courtesy output during tests
	LogOutput = ioutil.Discard
}

func TestEncDec(t *testing.T) {

	testset := []struct {
		key  string
		data string
	}{
		{key: "", data: "small data"},
		{key: "key", data: ""},
		{key: "utf8 key ∆∆", data: "utf8 data ∆∆"},
		{key: "key", data: strings.Repeat("large_data", 1<<10)},
		{key: strings.Repeat("large key", 1<<10), data: strings.Repeat("large data", 1<<10)},
		{key: strings.Repeat("large key", 1<<10), data: "sml data"},
	}

	for _, test := range testset {
		encdata, err := Encrypt([]byte(test.data), []byte(test.key))
		if err != nil {
			t.Error(err)
		}

		decdata, err := Decrypt(encdata, []byte(test.key))
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal([]byte(test.data), decdata) {
			t.Errorf("output mismatch")
		}
	}

}

func TestEncDecDirectory(t *testing.T) {

	const testSubDir = "test_enc_dec_directory"
	dir, err := ioutil.TempDir(os.TempDir(), testSubDir)
	if err != nil {
		t.Fatal(err)
	}

	// generate a test dir setup
	encryptTheseDir := filepath.Join(dir, "encrypt/these")
	if err := os.MkdirAll(encryptTheseDir, 0777); err != nil {
		t.Fatal(err)
	}

	dontEncryptTheseDir := filepath.Join(dir, "not/these")
	if err := os.MkdirAll(dontEncryptTheseDir, 0777); err != nil {
		t.Fatal(err)
	}

	const TextFile = "normal.txt"
	var TextFileData = bytes.Repeat([]byte("lines of data\n"), 500)
	var TextFileMode os.FileMode = 0644
	if err := ioutil.WriteFile(filepath.Join(encryptTheseDir, TextFile), []byte(TextFileData), TextFileMode); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(dontEncryptTheseDir, TextFile), []byte(TextFileData), TextFileMode); err != nil {
		t.Fatal(err)
	}

	const BinaryFile = "binary"
	var BinaryFileData = bytes.Repeat([]byte{0xff, 0xaa, 0x00}, 1<<10)
	var BinaryFileMode os.FileMode = 0755
	if err := ioutil.WriteFile(filepath.Join(encryptTheseDir, BinaryFile), []byte(BinaryFileData), BinaryFileMode); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(dontEncryptTheseDir, BinaryFile), []byte(BinaryFileData), BinaryFileMode); err != nil {
		t.Fatal(err)
	}

	const ReadOnlyFile = "read_only.🔒"
	var ReadOnlyFileData = bytes.Repeat([]byte(" 🔒 "), 1<<10)
	var ReadOnlyFileMode os.FileMode = 0400
	if err := ioutil.WriteFile(filepath.Join(encryptTheseDir, ReadOnlyFile), []byte(ReadOnlyFileData), ReadOnlyFileMode); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(dontEncryptTheseDir, ReadOnlyFile), []byte(ReadOnlyFileData), ReadOnlyFileMode); err != nil {
		t.Fatal(err)
	}

	var key = []byte("key")

	if err := EncryptDirFiles(encryptTheseDir, key); err != nil {
		t.Error(err)
	}

	var atLeastOneFileChecked bool

	if err := filepath.Walk(encryptTheseDir, func(path string, info os.FileInfo, err error) error {

		// skip processing directories as files
		if info.IsDir() {
			return nil
		}

		atLeastOneFileChecked = true

		if !sha256HexRegexp.MatchString(filepath.Base(path)) {
			return fmt.Errorf("Expected sha256 hash filename, got %q", filepath.Base(path))
		}

		// ensure we write with user read only perm
		if info.Mode().Perm()&0400 != 0400 {
			return fmt.Errorf("Expected %#o, got %#o", 0400, info.Mode().Perm())
		}

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

		return nil

	}); err != nil {
		t.Error(err)
	}

	if !atLeastOneFileChecked {
		t.Errorf("Somehow there were no files found in %q", dir)
	}

	// decrypt from top most temp dir and check to see if only encrypted files touched
	if err := DecryptDirFiles(dir, key); err != nil {
		t.Error(err)
	}

	atLeastOneFileChecked = false
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		// skip processing directories as files
		if info.IsDir() {
			return nil
		}

		atLeastOneFileChecked = true

		switch filepath.Base(path) {
		case TextFile:
			stat, err := os.Stat(path)
			if err != nil {
				return err
			}
			if stat.Mode().Perm() != TextFileMode {
				return fmt.Errorf("Expected decrypted file mode %#o, got %#o", TextFileMode, stat.Mode().Perm())
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if !bytes.Equal(data, TextFileData) {
				return fmt.Errorf("File data mismatch for %q", path)
			}

		case BinaryFile:
			stat, err := os.Stat(path)
			if err != nil {
				return err
			}
			if stat.Mode().Perm() != BinaryFileMode {
				return fmt.Errorf("Expected decrypted file mode %#o, got %#o", BinaryFileMode, stat.Mode().Perm())
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if !bytes.Equal(data, BinaryFileData) {
				return fmt.Errorf("File data mismatch for %q", path)
			}

		case ReadOnlyFile:
			stat, err := os.Stat(path)
			if err != nil {
				return err
			}
			if stat.Mode().Perm() != ReadOnlyFileMode {
				return fmt.Errorf("Expected decrypted file mode %#o, got %#o", ReadOnlyFileMode, stat.Mode().Perm())
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if !bytes.Equal(data, ReadOnlyFileData) {
				return fmt.Errorf("File data mismatch for %q", path)
			}

		default:
			return fmt.Errorf("Found unexpected file: %q", path)
		}

		return nil

	}); err != nil {
		t.Error(err)
	}

	if !atLeastOneFileChecked {
		t.Errorf("Somehow there were no files found in %q", dir)
	}

	if err := os.RemoveAll(dir); err != nil {
		t.Error(err)
	}

}
