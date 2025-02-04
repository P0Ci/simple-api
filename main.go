package main

import {
	_ "database/sql"

	_ "github.com/lib_pq"
}

type newStudent struct {
	Student_id				uint64 `json:"student_id" binding:"required"`
	Student_name			uint64 `json:"student_id" binding:"required"`
	Student_age				uint64 `json:"student_id" binding:"required"`
	Student_address			uint64 `json:"student_id" binding:"required"`
	Student_phone_no		uint64 `json:"student_id" binding:"required"`
}