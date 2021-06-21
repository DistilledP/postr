package util

func StringSlicePop(items []string) (string, []string) {
	return items[len(items)-1], items[:len(items)-1]
}
