package hgt

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
