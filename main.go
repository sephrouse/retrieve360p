package main

import (
  "fmt"
  "os"
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
  )

// syney, when you want to use this program, please remove all spaces in the following URL.
const GetPriceURL string = "h t t p s : / / p . 3 . c n / p r i c e s / g e t ?s k u I d = J _"

func getPriceFromSku(sku string) (price string, err error) {
  fullURL := GetPriceURL + sku
  resp, err := http.Get(fullURL)
  if err != nil {
    fmt.Println("getPriceFromSku : http error. err is ", err)
    return price, err
  }

  defer resp.Body.Close()

  var fields map[string]string

  contents, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("getPriceFromSku : ioutil readall error. err is ", err)
    return price, err
  }

  fmt.Println(string(contents))
  contentLen := len(contents)
  realContent := contents[1:contentLen - 2]
  fmt.Println(string(realContent))

  if err = json.Unmarshal(realContent, &fields); err != nil {
    fmt.Println("getPriceFromSku : json unmarshal error. err is ", err)
    return price, err
  }

  for k, v := range fields {
    fmt.Println(k, v)

    if k == "p" {
      price = v
    }
  }

  return price, err
}

func getSkuFromUrl(requestUrl string) (sku string, err error) {
  u, err := url.Parse(requestUrl)

  //fmt.Println(u.RequestURI())
  length := len(u.RequestURI())
  length = length - 5
  sku = (u.RequestURI())[1:length]

  return sku, err
}

func main() {
  //fmt.Println(len(os.Args))
  // syney, assume the os.Args[1] is absolutely a real item URL of jd.
  requestUrl := os.Args[1]

  sku, err := getSkuFromUrl(requestUrl)
  if err != nil {
    fmt.Println("main : cannot get sku from url. ", requestUrl)
    return
  }

  fmt.Println("main : sku is ", sku)

  price, err := getPriceFromSku(sku)
  if err != nil {
    return
  }

  fmt.Println("result : sku :", sku, ", price is ", price)
}
