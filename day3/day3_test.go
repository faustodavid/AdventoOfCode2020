package main

import "testing"

func Test_should_count_7_tree(t *testing.T) {
	inputData := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"
	grid := StringToByteGrid(inputData)
	treeCount := CountTreesOnGridUsingPattern(grid, 3, 1)

	if treeCount != 7 {
		t.Errorf("treeCount expected is 7 found %d", treeCount)
	}
}
