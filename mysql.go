package main

import (

    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

)

type Employes struct {
    EID   int    `json:"eid"`
    NAME string `json:"name"`
	EMAIL string `json:"email"`
}

var db *sql.DB


func connectDB() {

	dbb, err := sql.Open("mysql", "**********")

    if err != nil {
        panic(err.Error())        
	}

	db = dbb

	fmt.Println("Connected.......")
}

func insert() {

	insert, err := db.Query("insert into employes values(6,'Obito', 'obito@123')")

	if err != nil {

		panic(err.Error())
	}
	
	defer insert.Close()
	fmt.Println("Insert Done")
}

func read(eid int) {

	read, err := db.Query("select * from employes where eid = ?", eid)

	if err != nil {
		panic(err.Error())
	}
	
    for read.Next() {

        var emp Employes
        
		err = read.Scan(&emp.EID, &emp.NAME, &emp.EMAIL)
        
		if err != nil {
            panic(err.Error()) 
        }

		fmt.Println(emp.EID, emp.NAME, emp.EMAIL)
    }

	defer read.Close()
	fmt.Println("Read Done")
}

func delete() {

	delete, err := db.Query("delete from employes where eid = 3")

	if err != nil {
		panic(err.Error())
	}
	
	defer delete.Close()
	fmt.Println("Delete Done")
}

func update() {

	update, err := db.Query("update employes set name = 'changed' where eid = 2")

	if err != nil {
		panic(err.Error())
	}
	
	defer update.Close()
	fmt.Println("Update Done")
}

func main() {

    connectDB()
    defer db.Close()

	// insert()
	read(5)
    // delete()
	// update()
}

















// package main
 
// import (
//     "database/sql"
//     "fmt"
//     _ "github.com/go-sql-driver/mysql"
// )
 
// var dbm *sql.DB

// func connectDB() {
// 	db, err := sql.Open("mysql", "root:Tarun@/mydb")
    
// 	if err != nil {
//         panic(err.Error())
//     }

// 	defer db.Close()
// 	fmt.Println("connected")
// 	dbm  = db
// }

// func main() {

// 	connectDB()

// }
 
