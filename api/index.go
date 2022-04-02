package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
        "encoding/json"
)

type Response struct {
	Resultcode string `json:"resultcode"`
	Isvideo    int    `json:"isvideo"`
	Overstep   int    `json:"overstep"`
	URL        string `json:"url"`
}

func getFakeID() string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 22)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getUrl(ip, port string) *Response {
	host := "http://dir.wo186.tv:809/"
	path := "if5ax/?fakeid=" + getFakeID() + "&spid=31117&pid=31117&spip=" + ip + "&spport=" + port
	m := md5.Sum([]byte(path + "3d99ff138e1f41e931e58617e7d128e2"))
	key := hex.EncodeToString(m[:])

	resp, err := http.Get(host + path + "&spkey=" + key)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
  
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var result *Response
	if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
                fmt.Println("Can not unmarshal JSON")
	}
	return result
}

func main(w http.ResponseWriter, r *http.Request) {
	result := getUrl("23.224.47.131", "443")	
	fmt.Println(result)
}
