ackage main

import (
        "fmt"
        "bufio"
        "strings"
 //       "io"
        "log"
        "net/http"
        "os"
        "io/ioutil"
)

func get_url() {
        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
                os.Exit(1)
        }
        response, err := http.Get(os.Args[1])
        if err != nil {
                log.Fatal(err)
        } else {
                defer response.Body.Close()
                  body, err := ioutil.ReadAll(response.Body)
                  if err != nil { log.Fatal(err)  }
                  return string(body)
        }
}
func ReadFile() {

        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s <html_fime>\n", os.Args[0])
                os.Exit(1)
        }
         file, err := os.Open(os.Args[1])
         if err != nil { log.Fatal(err) }
          defer file.Close()
           scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        Line :=scanner.Text()
        switch{
              case strings.Contains(Line,"TCK test results smokes"):
                    Collection_name :="TCK"
              case strings.Contains(Line,"APPSERVER test results smokes"):
                    Collection_name :="APPSERVER"
              case strings.Contains(Line,"LRF ISV failure results"):
                    Collection_name :="ISV"
              case strings.Contains(Line,"LRF UNITTEST failure results"):
                    Collection_name :="UNITTEST"
              case strings.Contains(Line,"LRF APPSERVER failure results"):
                    Collection_name :="APPSERVER_LRF"
              case default:
                    break

              }
        if strings.Contains(Line,"not_reported"):
            //write to file


    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
        }

func main(){
s=get_url()
doc, err := html.Parse(strings.NewReader(s))
if err != nil {
    log.Fatal(err)
}
var f func(*html.Node)
f = func(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                fmt.Println(a.Val)
                break
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        f(c)
    }
}
f(doc)



}
