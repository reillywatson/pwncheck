package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pwncheck <mypassword>")
		os.Exit(1)
	}
	h := sha1.New()
	h.Write([]byte(os.Args[1]))
	hash := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	prefix := hash[:5]
	suffix := hash[5:]
	resp, err := http.Get(fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", prefix))
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), suffix) {
			fmt.Printf("Password pwned! Found %s times\n", strings.Split(scanner.Text(), ":")[1])
			os.Exit(1)
		}
	}
	fmt.Println("Password not pwned!")
}
