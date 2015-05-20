package main

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
)

func RentScrape() {
  doc, err := goquery.NewDocument("http://rj.bomnegocio.com/imoveis/aluguel/apartamentos?q=tijuca+-barra&ret=1040&sp=1") 
  if err != nil {
    log.Fatal(err)
  }

  jQuery('.col_2 a').each(function(){ $(this).attr('name') })

  doc.Find(".list_adsBN_item").Each(func(i int, s *goquery.Selection) {
    id := s.Find(".title mb5px a").Attr('name')
    name := s.Find(".title mb5px a").Text()
    link := s.Find(".title mb5px a").Attr('href')
    description := s.Find(".text mb5px detail_specific").Text()
    location := s.Find("i").Find('text mb5px detail_region').Text()
    fmt.Printf("Rent %d: %s - %s\n", i, band, title)
  })
}

func main() {
  RentScrape()
}