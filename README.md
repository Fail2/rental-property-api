# Rental Property API

A high-performance RESTful API built with Go and Beego (API Mode) that loads property data from a JSON source file into an in-memory store at startup and provides advanced filtering rules.

## Technical Requirements
- Go 1.22 or higher
- Beego v2 (://github.com)

## Setup & Installation

1. Clone the repository and navigate to the project directory:
   ```bash
   git clone https://github.com/Fail2/rental-property-api.git
   cd rental-property-api
   ```

2. Download and sync the required Go modules dependencies:
   ```bash
   go mod tidy
   ```

3. Ensure the source data file is placed inside the correct directory:
   - File location: `data/rental_properties.json`

## Running the Project & Swagger Generation

To run the application and ensure that the interactive Swagger documentation is correctly parsed via Beego's automated annotation engine, follow this exact execution sequence:

1. Generate the required comments router mapping file:
   ```bash
   bee generate routers
   ```

2. Clear the local Go build cache to avoid stale documentation metadata:
   ```bash
   go clean -cache
   ```

3. Start the application environment with the automated document generation flag enabled:
   ```bash
   bee run -gendoc=true
   ```
The server will initialize the in-memory data store, compile your annotations, and start listening on: `http://localhost:8080`

## Interactive API Documentation (Swagger UI)

Once the application environment is active, you can access the self-documenting live interactive dashboard using your web browser:
- **Swagger UI URL:** `http://localhost:8080/swagger`
- **Local API Specification Source:** If the dashboard shows default sample placeholders, clear the top input bar, paste `http://localhost:8080/swagger/swagger.json`, and click the green **Explore** button.

## Running Unit Tests

Execute the comprehensive table-driven service layer tests and handler assertions:
```bash
go test ./... -v
go vet ./...
```

## Sample cURL Commands

### 1. List Properties (Without Filters)
```bash
curl -X GET "http://localhost:8080/v1/properties"
```

### 2. List Properties with AND/OR Filtering & Pagination Limits
```bash
curl -X GET "http://localhost:8080/v1/properties?feed=11&published=false&min_price=50&max_price=250&property_type=Apartment&limit=3&amenities=Internet,Parking"
```

### 3. Get Single Property By Unique ID
```bash
curl -X GET "http://localhost:8080/v1/properties/BC-1000001"
```

### 4. Trigger Parameter Validation Error Response (400 Bad Request)
```bash
curl -X GET "http://localhost:8080/v1/properties?feed=abc"
```
