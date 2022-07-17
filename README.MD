# Installation
```bash
go mod tidy
```

# Database
```postgresql
     Database name : currency
     data table name : currencyrate
```

# Run Project
```bash
go run main.go
```
# End points
```bash
   - /rates/latest
        - Give latest rate difference between euro and other currency
   
   - /rates/{date} - Example: http://localhost:3000/rates/2022-06-15
        - Give rate difference of a specific date between euro and other currency
        
   - /rates/analyze 
        - Give minium , maximum and avarage between euro and other currency
   
```
# Test Case
```bash
    End point testing
        - go test -v routers/router_test.go routers/router.go 
```
# For WEB UI Test result
```bash
    Install goconvey 
      - go get github.com/smartystreets/goconvey
    
    In project folder command-line run
      - goconvey    
      
```
