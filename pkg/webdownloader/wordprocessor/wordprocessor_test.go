package wordprocessor

import (
	"strings"
	"testing"
)

func TestWordProcessor_UniqueWordsCount(t *testing.T) {
	testcases := []struct {
		input    string
		expected map[string]int
	}{
		{
			`September is a time
			Of beginning for all,
			Beginning of school
			Beginning of fall.`,
			map[string]int{
				"SEPTEMBER": 1,
				"IS":        1,
				"TIME":      1,
				"OF":        3,
				"BEGINNING": 3,
				"FOR":       1,
				"ALL":       1,
				"SCHOOL":    1,
				"FALL":      1,
			},
		},
		{
			`November. No shade, no shine,
			No_butterflies, no|bees,
			No=fruits, no-flowers,
			No^leaves, no+birds,
			November.`,
			map[string]int{
				"SHADE":       1,
				"SHINE":       1,
				"BUTTERFLIES": 1,
				"BEES":        1,
				"FRUITS":      1,
				"FLOWERS":     1,
				"LEAVES":      1,
				"BIRDS":       1,
				"NOVEMBER":    2,
				"NO":          8,
			},
		},
	}

	wp := New()
	for _, testcase := range testcases {
		res := wp.UniqueWordsCount(strings.NewReader(testcase.input))

		for k, v := range testcase.expected {
			if count, exist := res[k]; !exist {
				t.Fatalf("Missing result key %s", k)
			} else if count != v {
				t.Fatalf("Wrong result count for key %s. Want %d got %d", k, v, count)
			}
		}
	}
}
