package main

type AccountInfo struct {
    Id       string
    Name     string
    Lot      float32
    Lat      float32
    ShopName string
}

type ShopInfo struct {
    Name         string
    Lot          float32
    Lat          float32
    AccountCount int
}

type GeoItem struct {
    Name string
    Lot  float32 //经度
    Lat  float32 //纬度
}
