package grains

import "errors"

//Square returns the number of grains in a specific
//position in a chessboard
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid position")
	}

	return 1 << uint64(n-1), nil
}

// Total returns total number of grains in a chessboard
func Total() uint64 {
	return 1<<64 - 1
}
