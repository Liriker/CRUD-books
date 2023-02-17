# CRUD-books
### Service for interaction with book business entity
---
The service accepts JSON data and returns books in JSON format.

- GET    /books   --> Get all books
- GET    /book    --> Get book by id from JSON field
- POST   /book    --> Create book from JSON file
- PUT    /book    --> Update book from JSON fields by id
- DELETE /book    --> Delete book by id
---
The whole code is divided into 3 components:
- Endpoint. That component is responsible for interaction with the client. And at the moment, logging takes place in it.
- Service. That component is responsible for buiseneess logic and interaction between Endpoint and Repsitory.
- Repository. That component if responsible for interaction with data base.
---
During the development process, I mastered SOLID, trying to reproduce them in this work.
And this project is under development now.
