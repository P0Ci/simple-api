package main

import (
    "database/sql"
    "log"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

type newStudent struct {
    Student_id      uint64 `json:"student_id" binding:"required"`
    Student_name    string `json:"student_name" binding:"required"`
    Student_age     uint64 `json:"student_age" binding:"required"`
    Student_address string `json:"student_address" binding:"required"`
    Student_phone_no string `json:"student_phone_no" binding:"required"`
}

func setRouter() *gin.Engine {
    conn := "postgresql://postgres:postgre@127.0.0.1/postgres?sslmode=disable"
    
    db, err := sql.Open("postgres", conn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    r := gin.Default()
    return r
}

func main() {
    r := setRouter()
    r.Run(":5020")
}

