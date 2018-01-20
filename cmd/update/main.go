package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rudbast/tkpdld/config"
	"github.com/rudbast/tkpdld/ld"
	"github.com/rudbast/tkpdld/slack"
)

func init() {
	flag.Parse()
}

func main() {
	ctx := context.Background()

	err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = slack.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	count, err := ld.Get().CountLeft(ctx, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	err = slack.Get().UpdateStatus(ctx, reformatLD(count), ":kappa2:")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}

func reformatLD(count int) string {
	msg := fmt.Sprintf("It's H-%d to D-Day !!", count)
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
