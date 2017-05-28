package hgt

import (
	"bufio"
	"encoding/binary"
	"os"

	"github.com/pkg/errors"
)

// Load raw hgt file data into given tile
func LoadTileFromFile(sourceFile string, tile *Tile) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}

	err = binary.Read(bufio.NewReader(source), binary.BigEndian, tile.RawRow[:])
	if err != nil {
		return errors.Wrapf(err, "could not read tile data from %s", sourceFile)
	}

	return nil
}
