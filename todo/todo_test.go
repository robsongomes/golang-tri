package todo

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestItem(t *testing.T) {
	t.Run("TestLabel", func(t *testing.T) {
		var tests = []struct {
			input    Item
			expected string
		}{
			{Item{}, "0."},
			{Item{position: 1}, "1."},
		}
		for _, test := range tests {
			if result := test.input.Label(); result != test.expected {
				t.Errorf("item.Label() should be %s but got %s", test.expected, result)
			}
		}
	})

	t.Run("TestSetPriority", func(t *testing.T) {
		var tests = []struct {
			input    Item
			priority int
			result   int
		}{
			{Item{}, 0, 2},
			{Item{}, 1, 1},
			{Item{}, 2, 2},
			{Item{}, 3, 3},
			{Item{}, 4, 2},
		}

		for _, test := range tests {
			test.input.SetPriority(test.priority)
			if test.input.Priority != test.result {
				t.Errorf("SetPriority(%d) should result in priority=%d but result in priority=%d",
					test.priority, test.result, test.input.Priority)
			}
		}
	})

	t.Run("TestPrettyP", func(t *testing.T) {
		var tests = []struct {
			input  Item
			result string
		}{
			{Item{Priority: 1}, "(1)"},
			{Item{Priority: 2}, " "},
			{Item{Priority: 3}, "(3)"},
		}
		for _, test := range tests {
			result := test.input.PrettyP()
			if result != test.result {
				t.Errorf("PrettyP() with Priority=%d should return %s but received %s",
					test.input.Priority, test.result, result)
			}
		}
	})

	t.Run("TestPrettyDone", func(t *testing.T) {
		var tests = []struct {
			input  Item
			result string
		}{
			{Item{Done: true}, "X"},
			{Item{Done: false}, ""},
			{Item{}, ""},
		}
		for _, test := range tests {
			result := test.input.PrettyDone()
			if result != test.result {
				t.Errorf("PrettyDone() with Done=%v should return %s but received %s",
					test.input.Done, test.result, result)
			}
		}
	})
}

func TestSortItemsByPriority(t *testing.T) {
	t.Run("DoneFirst", func(t *testing.T) {
		var items = []Item{
			{Priority: 1},
			{Priority: 2},
			{Priority: 3, Done: true},
		}

		sort.Sort(ByPri(items))

		if items[0].Done != true {
			t.Errorf("First element of sorted Items should be Done")
		}
	})

	t.Run("HighPrioritiesFirst", func(t *testing.T) {
		var items = []Item{
			{Text: "I3", Priority: 3},
			{Text: "I2", Priority: 2},
			{Text: "I1", Priority: 1},
		}

		sort.Sort(ByPri(items))

		if items[0].Text != "I1" {
			t.Errorf("First element of sorted Items should be I1 but was %s", items[0].Text)
		}
	})

	t.Run("PositionAndPriority", func(t *testing.T) {
		var items = []Item{
			{Text: "I3", Priority: 3, position: 3},
			{Text: "I2", Priority: 3, position: 2},
			{Text: "I1", Priority: 3, position: 1},
		}

		sort.Sort(ByPri(items))

		if items[0].Text != "I1" {
			t.Errorf("First element of sorted Items should be I1 but was %s", items[0].Text)
		}
	})
}

func TestSaveItems(t *testing.T) {
	items := []Item{
		{Text: "Item 01", position: 1, Priority: 3, Done: false},
		{Text: "Item 02", position: 2, Priority: 2, Done: true},
	}
	jsonExpected, _ := json.Marshal(items)

	saved := writeInFile
	defer func() { writeInFile = saved }()

	var jsonResult []byte

	writeInFile = func(filename string, b []byte) error {
		jsonResult = b
		return nil
	}

	SaveItems("teste.txt", items)

	if string(jsonExpected) != string(jsonResult) {
		t.Errorf("Content write in file is different. Expected: \n%s \ngot \n%s", jsonExpected, jsonResult)
	}
}

func TestReadItems(t *testing.T) {
	bytes := []byte(`[{"Text":"Item 01","Priority":3,"Done":false},{"Text":"Item 02","Priority":2,"Done":true}]`)

	saved := readFile
	defer func() { readFile = saved }()

	var items []Item

	readFile = func(filename string) ([]byte, error) {
		return bytes, nil
	}

	items, _ = ReadItems("teste.txt")

	if len(items) != 2 {
		t.Errorf("The len of items read should be %d but was %d", 2, len(items))
	}

	if items[0].Text != "Item 01" {
		t.Errorf("First Item should be %s but was %s", "Item 01", items[0].Text)
	}
}
