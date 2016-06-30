package main

import (
    "bufio"
    "errors"
    _ "github.com/mattn/go-sqlite3"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
)

var accountJsFile = "assets/datajs/account_data.js"
var shopJsFile = "assets/datajs/shop_data.js"

var shopMap map[string]ShopInfo = make(map[string]ShopInfo)
var accountMap map[string]AccountInfo = make(map[string]AccountInfo)

func GenerateJs(dataFiles []string) {
    log.Println("parse file:", dataFiles)
    for _, f := range dataFiles {
        parseDataFile(f)
    }
    generateDataJs()
}

func ImportDb(dataFiles []string, dataService DataService) {
    for _, f := range dataFiles {
        parseDataFile(f)
    }

    dataService.RemoveDB()
    err := dataService.CreateAccountTable()
    if err != nil {
        log.Fatal(err)
    }

    i := 0
    flag := 10
    for _, acc := range accountMap {
        i ++
        if i == flag {
            log.Println("已导入数据条数：", flag)
            flag = flag * 10
        }
        dataService.SaveAccount(&acc)
    }
    log.Println("帐号数据导入完成。总条数：", len(accountMap))
}

func parseDataFile(dataFile string) {
    log.Println("parse file:", dataFile)
    df, err := os.Open(dataFile)
    if err != nil {
        log.Fatalln(err)
        return
    }
    defer df.Close()

    var totalLine, errorLine int
    scanner := bufio.NewScanner(df)
    for scanner.Scan() {
        line := scanner.Text()
        totalLine++
        err = parseLine(line)
        if err != nil {
            log.Println(err)
            errorLine++
        }
    }

    log.Println("total line:", strconv.Itoa(totalLine))
    log.Println("error line:", strconv.Itoa(errorLine))
}

func strToFloat32(s string) float32 {
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        return float32(0.0)
    } else {
        return float32(f)
    }
}

func parseLine(line string) error {
    flds := strings.Split(line, "\t")
    log.Println("flds:", strings.Join(flds, " , "))
    if len(flds) < 15 {
        log.Println("invalid fld len:", strconv.Itoa(len(flds)))
        return errors.New("invalid data: " + line)
    }
    shopName := flds[9]
    shopLa := strToFloat32(flds[10])
    shopLo := strToFloat32(flds[11])
    if shopInfo, ok := shopMap[shopName]; ok {
        shopInfo.AccountCount = shopInfo.AccountCount + 1
        shopMap[shopName] = shopInfo
    } else {
        newShopInfo := ShopInfo{Name: shopName, Lat: shopLa, Lot: shopLo, AccountCount: 1}
        shopMap[shopName] = newShopInfo
    }

    accId := flds[0]
    accName := flds[1]
    accLa := strToFloat32(flds[5])
    accLo := strToFloat32(flds[6])
    accInfo := AccountInfo{accId, accName, accLa, accLo, shopName}
    accountMap[accId] = accInfo
    return nil
}

func generateDataJs() {
    log.Println("shop size:", len(shopMap))
    log.Println("account size:", len(accountMap))
    shopJs := "var shopLocs = ["
    for shopName, shopInfo := range shopMap {
        // if !strings.HasPrefix(shopName, "沈阳") {
        //     continue;
        // }

        shopLoc := "{name:'" + shopName + "', value: " +
            strconv.Itoa(shopInfo.AccountCount) + ", geoCoord:[" +
            strconv.FormatFloat(float64(shopInfo.Lat), 'f', 6, 32) + ", " +
            strconv.FormatFloat(float64(shopInfo.Lot), 'f', 6, 32) + "]},"
        shopJs = shopJs + "\n" + shopLoc
    }
    shopJs = shopJs + "\n];"

    accountJs := "var accountLocs = ["
    for _, accInfo := range accountMap {
        accLoc := "{name:'" + accInfo.Name + "', value:1, geoCoord:[" +
            strconv.FormatFloat(float64(accInfo.Lat), 'f', 6, 32) + ", " +
            strconv.FormatFloat(float64(accInfo.Lot), 'f', 6, 32) + "]},"
        accountJs = accountJs + "\n" + accLoc
    }
    accountJs = accountJs + "\n];"

    ioutil.WriteFile(shopJsFile, []byte(shopJs), 0600)
    ioutil.WriteFile(accountJsFile, []byte(accountJs), 0600)
}
