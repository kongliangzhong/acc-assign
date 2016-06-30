package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
)

type DataService interface {
    GetAllAccount() []AccountInfo
    GetAllShop() []ShopInfo
    RemoveDB()
    CreateAccountTable() error
    CreateShopTable() error
    SaveAccount(account *AccountInfo)
}

type DataServiceImplSqlite3 struct {
    DBFile string
    Db     *sql.DB
}

var queryAllAccountSql = "select id, name, lot, lat, shopname from account"
func (ds *DataServiceImplSqlite3) GetAllAccount() []AccountInfo {
    rows, err := ds.Db.Query(queryAllAccountSql)
    if err != nil {
        log.Println("query account error:", err)
        return nil
    }
    defer rows.Close()
    allAccount := make([]AccountInfo, 0, 100)
    for rows.Next() {
        var id string
        var name string
        var lot float64
        var lat float64
        var shopname string
        err = rows.Scan(&id, &name, &lot, &lat, &shopname)
        if err != nil {
            log.Println("parse row error:", err, rows)
            continue
        }
        accountInfo := AccountInfo{id, name, float32(lot), float32(lat), shopname}
        allAccount = append(allAccount, accountInfo)
    }
    err = rows.Err()
    if err != nil {
        log.Println("query or parse error:", err)
    }

    return allAccount
}

func (ds *DataServiceImplSqlite3) GetAllShop() []ShopInfo {
    return nil
}

func (ds *DataServiceImplSqlite3) RemoveDB() {
    os.Remove(ds.DBFile)
}

func (ds *DataServiceImplSqlite3) CreateAccountTable() error {
    createAccountTableSql := `create table account(
      id string not null primary key,
      name string,
      lot   float,
      lat   float,
      shopname string
    )`
    _, err := ds.Db.Exec(createAccountTableSql)
    return err
}

func (ds *DataServiceImplSqlite3) CreateShopTable() error {
    return nil
}

var saveAccountSql = "insert into account(id, name, lot, lat, shopname) values(?, ?, ?, ?, ?)"

func (ds *DataServiceImplSqlite3) SaveAccount(account *AccountInfo) {
    stmt, err := ds.Db.Prepare(saveAccountSql)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(account.Id, account.Name, account.Lot, account.Lat, account.ShopName)
    if err != nil {
        log.Fatal(err)
    }
}
