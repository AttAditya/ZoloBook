# ZoloBooks Backend

This repository is made for the selection process of ZoloStays. It contains the backend for the ZoloBooks application.

> **ZoloBooks** is a simple application that allows users to share books with each other. Users can add books to their library and other users can request to borrow books from them. The application also allows users to search for books and users.

The backend is written in [Go](https://golang.org/) and uses the [Gin](https://github.com/gin-gonic/gin) framework.

> **Note:** This repository is not meant to be used in production. It is only for the selection process.

## API Usage

### GET `/api/v1/books`

> since: 0.0.1, last updated: 0.0.2

Returns a map of all books in the database.
For each book, the map contains the following fields:

- `name`: The name of the book
- `owner`: The id of the user who owns the book

Example:

```json
{
    "B001": {
        "name": "Book 1",
        "owner": "U001",
        "status": "AVAILABLE"
    },
    "B002": {
        "name": "Book 2",
        "owner": "U002",
        "status": "AVAILABLE"
    },
    "B003": {
        "name": "Book 3",
        "owner": "U003",
        "status": "AVAILABLE AFTER 2020-01-03 00:00:00"
    }
}
```

### GET `/api/v1/books/:id`

> since: 0.0.2

Returns the details of the book with the given id.

Example:

```json
{
    "name": "Book 1",
    "owner": "U001",
    "status": "AVAILABLE"
}
```

### PUT `/api/v1/books`

> since: 0.0.2

Adds a new book to the database.

Example request body:

```json
{
    "name": "Book 1",
    "owner": "U001"
}
```

### GET `/api/v1/books/borrowed`

> since: 0.0.2

Returns a map of all books borrowed by users.

Example:

```json
{
    {
        "R001": {
            "book_id": "B003",
            "borrow_end_time": "2020-01-03 00:00:00",
            "borrow_start_time": "2020-01-01 00:00:00",
            "borrower": "U004"
        },
        "R002": {
            "book_id": "B001",
            "borrow_end_time": "2020-01-03 00:00:00",
            "borrow_start_time": "2020-01-01 00:00:00",
            "borrower": "U005"
        }
    },
}
```

### GET `/api/v1/books/borrowed/:id`

> since: 0.0.2

Returns the details of the book borrowed by the user with the given id.

Example:

```json
{
    "book_id": "B003",
    "borrow_end_time": "2020-01-03 00:00:00",
    "borrow_start_time": "2020-01-01 00:00:00",
    "borrower": "U004"
}
```

### PUT `/api/v1/books/:id/borrow`

> since: 0.0.2

Borrows the book with the given id.

Example request body:

```json
{
    "borrower": "U004",
    "borrow_start_time": "2020-01-01 00:00:00",
    "borrow_end_time": "2020-01-03 00:00:00"
}
```

### POST `/api/v1/books/borrowed/:id/return`

> since: 0.0.2

Returns the book borrowed by the user with the given id.

Example request body:

```json
{
    "borrow_end_time": "2020-01-03 00:00:00"
}
```

