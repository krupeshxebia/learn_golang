### Learning Golang
---
#### Go Crash Course #1
- Hello World
- Variable & Types
- Packages
- Functions
- Arrays / Slices
- Conditionals
- Loops
- Maps
- Range
- Pointers
- Structs
- Interfaces
- Web

#### CRUD REST API

- **GET All Books**
GET : `http://localhost:8000/api/books`

- **Get Single Book From ID**
GET : `http://localhost:8000/api/books/{id}`

- **Add New Book**
POST      : `http://localhost:8000/api/books`
JSON Body : 
```json
{
    "isbn" : "123456",
    "title" : "Book Three",
    "author" : {
        "firstname" : "Marshall",
        "lastname" : "Williams"
    }
}
```
 
- **Delete A Book From ID**
DELETE : `http://localhost:8000/api/books/{id}`

- **Update A Book From ID**
PUT       : `http://localhost:8000/api/books/{id}`
JSON Body : 
```json
{
    "isbn" : "123456",
    "title" : "Updated Title",
    "author" : {
        "firstname" : "Jeff",
        "lastname" : "Thomson"
    }
}
```
 
 
 
 
