package handler

import (
  "net/http"
  "fmt"
  "math/rand"
  "time"
  "crypto/md5"
  "encoding/hex"
  "io"
  "encoding/json"
  "strings"
)

func getFakeID() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strb := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result []byte
	for i := 0; i < 22; i++ {
		result = append(result, strb[r.Intn(len(strb))])
	}
	return string(result)
}
func getUrl(ip, port string) []string {
	path := "if5ax/?fakeid=" + getFakeID() + "&spid=81117&pid=81117&spip=" + ip + "&spport=" + port
	host := "http://dir.wo186.tv:809/"
	m := md5.Sum([]byte(path + "3d99ff138e1f41e931e58617e7d128e2"))
	key := hex.EncodeToString(m[:])
	r, _ := http.Get(host + path + "&spkey=" + key)
	body, _ := io.ReadAll(r.Body)
	rj := map[string]string{}
	json.Unmarshal(body, &rj)
	p := strings.Index(rj["url"], "/if5ax")
	t := strings.Index(rj["url"], "lsttm=")
	return []string{rj["url"][7:p], rj["url"][p:], rj["url"][t+6 : t+20]}
}
func main(){
  url:=getUrl("1.1.1.1", "443")
  fmt.Println(url)
}

func Handler(w http.ResponseWriter, r *http.Request) {
  main()
}
