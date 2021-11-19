// Package oshash provides a library and a command-line interface tool for
// computing OpenSubtitles hash values for files and arbitrary data.
//
// The OpenSubtitles hash format is defined as a 64-bit integer value, made from
// the size of the input data in bytes, a checksum of the first 64 kibibytes in
// the data (the head), and a checksum of the last 64 kibibytes in the data
// (the tail).
//
// A notable limitation is that a piece of data must be at least 64 kibibytes
// long to have a valid OpenSubtitles hash. Supplying a file or []byte of less
// than 64 kibibytes will result in the error oshash.ErrDataTooSmall.
//
// All hash values returned by these functions are uint64s, but they can be
// trivially converted to hexadecimal strings with
// strconv.FormatUint(hash_value, 16).
package oshash

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

const ChunkSize = 65536

var ErrDataTooSmall = errors.New("data is less than 65536 bytes")
var ErrReadTooFewBytes = fmt.Errorf("failed to read %v bytes", ChunkSize)

// FromFile computes the oshash for an *os.File.
func FromFile(file *os.File) (uint64, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	if fi.Size() < ChunkSize {
		return 0, ErrDataTooSmall
	}
	buf := make([]byte, ChunkSize*2)
	if err = readChunk(file, 0, buf[:ChunkSize]); err != nil {
		return 0, err
	}
	if err = readChunk(file, fi.Size()-ChunkSize, buf[ChunkSize:]); err != nil {
		return 0, err
	}
	return computeOsHash(buf, uint64(fi.Size()))
}

// FromFilepath computes the oshash of the specified file.
func FromFilepath(filename string) (hash uint64, err error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return FromFile(f)
}

// FromFilepath computes the oshash of some given data.
func FromBytes(b []byte) (hash uint64, err error) {
	if len(b) < ChunkSize {
		return 0, ErrDataTooSmall
	}
	buf := make([]byte, ChunkSize*2)
	copy(buf[:ChunkSize], b[0:ChunkSize])
	copy(buf[ChunkSize:], b[(len(b)-ChunkSize):])
	return computeOsHash(buf, uint64(len(b)))
}

// readChunk reads file into buf starting at offset, and returns an error if it
// did not read at least 64 kibibytes.
func readChunk(file *os.File, offset int64, buf []byte) (err error) {
	n, err := file.ReadAt(buf, offset)
	if err == nil && n != ChunkSize {
		return ErrReadTooFewBytes
	}
	return err
}

// computeOsHash computes a final oshash value from buf, a bytearray of the
// sampled data.
func computeOsHash(buf []byte, filesize uint64) (hash uint64, err error) {
	var nums [(ChunkSize * 2) / 8]uint64
	reader := bytes.NewReader(buf)
	if err = binary.Read(reader, binary.LittleEndian, &nums); err != nil {
		return 0, err
	}
	for _, num := range nums {
		hash += num
	}
	return hash + filesize, nil
}
