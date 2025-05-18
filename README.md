# FIO Service

This project is a RESTful information service for processing full name data. It automatically enriches the input with estimated age, gender, and nationality using public APIs, and stores the results in a PostgreSQL database.

The system provides the following functionalities for users:
- Adding, updating, and deleting persons via REST API
- Receiving and enriching personal data via external APIs (agify.io, genderize.io, nationalize.io)
- Saving enriched information to the database
- Retrieving stored data with filters and pagination
- Swagger documentation for API exploration

Stack: Go, Gin, PostgreSQL, GORM, Swagger, Golang-Migrate, Logrus