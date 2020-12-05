package internal

var (
	possibleRows = getPossibleRows()
	possibleColumns = getPossibleColumns()
)

type BoardingPass struct {
	Row int
	Column int
}

func (b *BoardingPass) SeatID() int {
	return b.Row * 8 + b.Column
}

func NewBoardingPass(code string) *BoardingPass {
	return &BoardingPass{
		Row: decodeRow(code[0:7]),
		Column: decodeColumn(code[7:]),
	}
}

func decodeRow(rowCode string) int {
	return decode(rowCode, possibleRows, "B")
}

func decodeColumn(columnCode string) int {
	return decode(columnCode, possibleColumns, "R")
}

func decode(code string, values []int, upChar string) int {
	index := 0
	upper := len(values) - 1
	lower := 0
	for codeIndex, c := range code {
		if codeIndex == len(code) - 1 {
			if string(c) == upChar {
				index = upper
			} else {
				index = lower
			}
			break
		}

		index = ((upper + (upper % 2)) - lower) / 2 + lower
		if string(c) == upChar {
			lower = index
		} else {
			//lower = upper
			upper = index - (upper % 2)
		}
	}

	return values[index]
}

func getPossibleRows() []int {
	possibleRows := make([]int, 0)
	for i:=0; i < 128; i++ {
		possibleRows = append(possibleRows, i)
	}
	return possibleRows
}

func getPossibleColumns() []int {
	possibleColumns := make([]int, 0)
	for i:=0; i < 8; i++ {
		possibleColumns = append(possibleColumns, i)
	}
	return possibleColumns
}