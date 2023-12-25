# Description: This file contains curl commands to test the API endpoints

### End-points ###

## get books
#curl -X GET http://localhost:8000/api/v1/books
#
## get book by id
#curl -X GET http://localhost:8000/api/v1/books/B001
#
## put book
#curl -X PUT -H "Content-Type: application/json" -d '{
#    "name": "Book 4",
#    "user": "U004"
#}' http://localhost:8000/api/v1/books
#
## existing requests for all books
#curl -X GET http://localhost:8000/api/v1/books/borrowed
#
## request book
#curl -X PUT -H "Content-Type: application/json" -d '{
#    "borrower": "U005",
#    "borrow_start_time": "2020-01-01 00:00:00",
#    "borrow_end_time": "2020-01-03 00:00:00"
#}' http://localhost:8000/api/v1/books/B001/borrow
#
## return book
#curl -X POST -H http://localhost:8000/api/v1/books/borrowed/R002/return

### Tests ###

# 1: put book test
curl -X GET http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books/B001
echo ""
curl -X PUT -H "Content-Type: application/json" -d '{
    "name": "Book 4",
    "user": "U004"
}' http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books/B004
echo ""
echo ""

# 2: request book test
curl -X GET http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books/borrowed
echo ""
curl -X PUT -H "Content-Type: application/json" -d '{
    "borrower": "U005",
    "borrow_start_time": "2020-01-01 00:00:00",
    "borrow_end_time": "2020-01-03 00:00:00"
}' http://localhost:8000/api/v1/books/B001/borrow
echo ""
curl -X GET http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books/borrowed
echo ""
curl -X GET http://localhost:8000/api/v1/books/borrowed/R002
echo ""
echo ""

# 3: return book test
curl -X GET http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books/borrowed
echo ""
curl -X POST http://localhost:8000/api/v1/books/borrowed/R002/return
echo ""
curl -X GET http://localhost:8000/api/v1/books
echo ""
curl -X GET http://localhost:8000/api/v1/books/borrowed
echo ""
echo ""

