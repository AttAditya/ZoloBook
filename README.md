# ZoloBooks Backend

This repository is made for the selection process of ZoloStays. It contains the backend for the ZoloBooks application.

> **ZoloBooks** is a simple application that allows users to share books with each other. Users can add books to their library and other users can request to borrow books from them. The application also allows users to search for books and users.

The backend is written in [Go](https://golang.org/) and uses the [Gin](https://github.com/gin-gonic/gin) framework.

> **Note:** This repository is not meant to be used in production. It is only for the selection process.

## API Usage

### GET `/api/v1/books`

> since: 0.0.1

Returns a map of all books in the database.
For each book, the map contains the following fields:

- `name`: The name of the book
- `owner`: The id of the user who owns the book

Example:

```json
{
    "B001": {
        "name": "Book 1",
        "owner": "U001"
    },
    "B002": {
        "name": "Book 2",
        "owner": "U002"
    },
    "B003": {
        "name": "Book 3",
        "owner": "U003"
    }
}
```

