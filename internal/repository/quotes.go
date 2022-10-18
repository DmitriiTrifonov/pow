package repository

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Quotes struct {
	data []string
}

func NewQuotes(path string) (*Quotes, error) {
	path = filepath.Clean(path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	quotes := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.SplitN(line, " ", 2)
		if len(lines) > 1 {
			quotes = append(quotes, lines[1])
		}
	}
	return &Quotes{data: quotes}, nil
}

func (q *Quotes) GetRandomQuote() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(len(q.data))
	return q.data[num]
}
