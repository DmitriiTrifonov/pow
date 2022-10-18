package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"

	"github.com/DmitriiTrifonov/pow/internal/repository"
)

type RandomQuoter interface {
	GetRandomQuote() string
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	repo, err := repository.NewQuotes("text/word_of_wisdom.txt")
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot init Word of Wisdom repo"))
	}

	// TODO: Add cancel func
	ctx, _ := context.WithCancel(context.Background())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// TODO: Add error handling
		go handleTCPConn(ctx, conn, repo)
	}
}

func handleTCPConn(ctx context.Context, conn net.Conn, quoter RandomQuoter) error {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for loop := true; loop; {
		select {
		case <-ctx.Done():
			loop = false
		default:
			msg, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					loop = false
					break
				}
				log.Println(err)
				break
			}

			if msg != "" {
				log.Println(msg[:len(msg)-1])
			}

			_, err = writer.WriteString(quoter.GetRandomQuote() + "\n")
			err = writer.Flush()
			if err != nil && err != io.EOF {
				log.Println(err)
				continue
			}
		}
	}
	return conn.Close()
}
