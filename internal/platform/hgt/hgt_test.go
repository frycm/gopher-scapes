package hgt_test

import (
	"os"
	"testing"

	"github.com/frycm/gopher-scapes/internal/platform/hgt"
)

const TestFileName = "../../../test_data/raw_hgt/N50E015.hgt"

// Try to load N50E15 tile and check if max height is as expected
// Max height should be Snezka mountain (1603 m, 1600 in SRTMGL1)
func TestLoadTileFromFile(t *testing.T) {
	var tile hgt.Tile
	source, err := os.Open(TestFileName)
	if err != nil {
		t.Fatalf("Could open source file source file: %s", err)
	}
	err = hgt.Load(&tile, source)

	if len(tile.RawRow) != hgt.TileRawRowSize {
		t.Errorf("%d height point was expected in tile, but %d found", hgt.TileRawRowSize, len(tile.RawRow))
	}

	maxHeight := struct {
		row    int
		column int
		count  int
		value  int16
	}{}

	for i, height := range tile.RawRow {
		if height >= maxHeight.value {
			maxHeight.row = i / hgt.EdgeSize
			maxHeight.column = i % hgt.EdgeSize
			if height == maxHeight.value {
				maxHeight.count += 1
			} else {
				maxHeight.count = 1
				maxHeight.value = height
			}
		}
	}

	if maxHeight.row != 950 {
		t.Errorf("Max height row was expected %s, but was %s", 950, maxHeight.row)
	}
	if maxHeight.column != 2664 {
		t.Errorf("Max height column was expected %s, but was %s", 2664, maxHeight.column)
	}
	if maxHeight.count != 1 {
		t.Errorf("Max height count was expected %s, but was %s", 1, maxHeight.count)
	}
	if maxHeight.value != 1600 {
		t.Errorf("Max height was expected %s, but was %s", 1600, maxHeight.value)
	}
}
