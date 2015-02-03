package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandSeq(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz01234567ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	i, l := 0, len(chars)
	b := make([]byte, n)
	for i < n {
		b[i] = chars[rand.Intn(l)]
		i++
	}
	return string(b)
}

func JSON(w http.ResponseWriter, v interface{}, c int) {
	b, _ := json.MarshalIndent(v, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	w.Write(b)
}

func GetSegs(r *http.Request) []string {
	return strings.Split(strings.Trim(r.URL.Path, "/"), "/")
}

func GetSeg(r *http.Request, n int) string {
	segs := GetSegs(r)

	if len(segs) >= n {
		return segs[n-1]
	}

	return ""
}
