package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nytimes/settarino/catalog"
	"github.com/nytimes/settarino/sets"
)

const (
	iOSAlpha     = "0123456789abcdef"
	AndroidAlpha = "-_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	installtype string
	seglabel    string
	count       int
)

func init() {
	flag.StringVar(&installtype, "t", "", "installation type")
	flag.IntVar(&count, "c", 0, "number of tokens to generate")
	flag.StringVar(&seglabel, "l", "", "the segment label to use for this data")
}

func main() {
	flag.Parse()

	alpha, bits := iOSAlpha, 64
	switch installtype {
	case "newsfusion":
	case "newsandroid":
		alpha, bits = AndroidAlpha, 155
	case "newsiphone":
	case "crosswordsios":
	case "crosswordsandroid":
		alpha, bits = AndroidAlpha, 155
	}

	keys := make([]string, count, count)
	for i := 0; i < count; i++ {
		keys[i] = RandomToken(bits, alpha)
	}

	pset := sets.NewPrimitiveSet(time.Now(), sets.CanonicalTag(seglabel), keys)
	name := catalog.Name(installtype, pset)

	fo, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	fw := bufio.NewWriter(fo)
	if err := catalog.PersistSet(fw, pset); err != nil {
		log.Fatal(err)
	}
}

func RandomToken(n int, alpha string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alpha[rand.Int63()%int64(len(alpha))]
	}
	return string(b)
}
