package main

import (
  "fmt"
  "log"
  "strings"
  "github.com/PuerkitoBio/goquery"
)

func RentScrape() {
  doc, err := goquery.NewDocument("http://rj.olx.com.br/rio-de-janeiro-e-regiao/zona-norte/imoveis/aluguel/apartamentos?ret=1040") 
  if err != nil {
    log.Fatal(err)
  }

  doc.Find(".section_OLXad-list ul.list li").Each(func(i int, s *goquery.Selection) {
    anchor := s.Find("a")
    id := anchor.AttrOr("name","")[0]
    name := scrub(anchor.Text())
    // link := anchor.AttrOr("href","")
    // description := s.Find(".detail_specific").Text()
    price := scrub(s.Find(".OLXad-list-price").Text())
    // location := s.Find("i").Find(".detail_region").Text()
    fmt.Printf("%v\n", id)
    fmt.Printf("%v\n", name)
    fmt.Printf("%v\n", price)
  })
}

func main() {
  RentScrape()
}

func scrub(s string) string{
  return strings.Join( strings.Fields(s), " ") 
}