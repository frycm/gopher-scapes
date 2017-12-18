package hgt

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
)

const (
	// Tile edge size
	EdgeSize = 3601
	// Flattened raw row size (rows x cols)
	TileRawRowSize = EdgeSize * EdgeSize
	// Point containing missing data
	MissingData = -32768
)

// 1 arc second height data wrapper
type Tile struct {
	// Height data flattened into one dimensional array
	RawRow [TileRawRowSize]int16
}

// Method for obtaining point data in conventional way [row,col]
func (tile *Tile) Point(row uint, col uint) int16 {
	if tile == nil {
		return MissingData
	}
	return tile.RawRow[row*EdgeSize+col]
}

// Load raw hgt data into given source
func Load(targetTile *Tile, source io.Reader) error {
	err := binary.Read(source, binary.BigEndian, targetTile.RawRow[:])
	if err != nil {
		return errors.Wrapf(err, "could not read tile data from source")
	}

	return nil
}
