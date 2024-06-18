package utils

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid/v2"
)

type Ulid string

func (u Ulid) String() string {
	return string(u)
}

func NewUlid() Ulid {
	t := time.Now()
	entropy := ulid.Monotonic(ulid.DefaultEntropy(), 0)
	return Ulid(ulid.MustNew(ulid.Timestamp(t), entropy).String())
}

func CloneRequest(r *http.Request) *http.Request {
	var bodyBytes []byte
	newRequest := *r.WithContext(r.Context())

	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	newRequest.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return &newRequest
}

func RandomElementFromSlice[T any](items []T) T {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a random index
	randomIndex := r.Intn(len(items))
	return items[randomIndex]
}
