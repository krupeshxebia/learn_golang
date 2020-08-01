## Learning Golang

### 📁 Go Crash Course #1
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

### 📁 CRUD REST API #1

**GET All Books**<br>
GET : `http://localhost:8000/api/books`

**Get Single Book From ID**<br>
GET : `http://localhost:8000/api/books/{id}`

**Add New Book**<br><br>
POST      : `http://localhost:8000/api/books`<br>
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
 
**Delete A Book From ID**<br>
DELETE : `http://localhost:8000/api/books/{id}`

**Update A Book From ID**<br>
PUT       : `http://localhost:8000/api/books/{id}`<br>
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
 
### 📁 REST API #2
 
* [] `GET /coasters` returns list of coasters as JSON
* [] `GET /coasters/{id}` returns details of specific coaster as JSON
* [] `POST /coasters` accepts a new coaster to be added
* [] `POST /coasters` returns status 415 if content is not `application/json`
* [] `GET /admin` requires basic auth
* [] `GET /coasters/random` redirects (Status 302) to a random coaster
 
 
