package assign1

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercolation(t *testing.T) {
	tests := []struct {
		name string
		file string
		want bool
	}{
		{"input2", "./input/input2.txt", true},
		{"input2-no", "./input/input2-no.txt", false},
		{"input3", "./input/input3.txt", true},
		{"input4", "./input/input4.txt", true},
		{"input5", "./input/input5.txt", true},
		{"input6", "./input/input6.txt", true},
		{"input7", "./input/input7.txt", true},
		{"input8", "./input/input8.txt", true},
		{"input8-no", "./input/input8-no.txt", false},
		{"input10", "./input/input10.txt", true},
		{"input10-no", "./input/input10-no.txt", false},
		{"input20", "./input/input20.txt", true},
		{"input50", "./input/input50.txt", true},
		{"java60", "./input/java60.txt", true},
		{"jerry47", "./input/jerry47.txt", true},
		{"princeton96", "./input/princeton96.txt", true},
		{"sedgewick60", "./input/sedgewick60.txt", true},
		{"snake13", "./input/snake13.txt", true},
		{"snake101", "./input/snake101.txt", true},
		{"wayne98", "./input/wayne98.txt", true},
	}

	var size int
	for _, v := range tests {
		f, err := os.Open(v.file)

		if err == nil {
			_, err = fmt.Fscanf(f, "%d\n", &size)
			t.Log(v.name, err, size)
		}

		p := NewPercolation(size)

		var row, col int
		for {
			if err != nil {
				break
			}
			_, err = fmt.Fscanf(f, "%d %d\n", &row, &col)
			if err == nil {
				t.Log(v.name, "open: ", row, col)
				p.Open(row, col)
				t.Log(v.name, p.LogData())
			} else {
				t.Log(v.name, "Err: ", err)
			}
		}

		got := p.Percolates()
		t.Log(v.name, p.LogData(), p.uf.LogData())
		assert.Equal(t, v.want, got)
	}
}
