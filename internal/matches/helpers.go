package matches

import "strconv"

func score(score *int) string {
	if score == nil {
		return "-"
	}

	return strconv.Itoa(*score)
}
