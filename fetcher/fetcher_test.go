package fetcher

import "testing"

func TestFetch(t *testing.T) {
	var url = "https://www.pilishen.com/posts"
	result, err := Fetch(url)
	if err != nil {
		t.Fatalf("fetch err: %v", err)
	}
	t.Logf("html result:%s", result)
}