<h1 align="center">Ubiwhere Challenge - Elisio Pinheiro</h1>


---

## 📝 Table of Contents

- [Getting Started](#getting_started)
- [CLI](#commands)
- [REST](#rest)
- [Project Structure](#project_structure)
- [Acknowledgments](#acknowledgement)

## ⛏️ Built Using
- [Golang](https://golang.org/) - v1.14 (Required)
    - [Cobra](https://github.com/spf13/cobra)
    - [GORM](https://gorm.io/)
    - [Gin](https://github.com/gin-gonic/gin)


## 🏁 Getting Started <a name = "getting_started"></a>

### Get the project

You can place the project wherever you want, using GOMODULES enabled.

```
git clone https://github.com/elisiopinheiro/ubiwherechallenge.git
```

### Installing

Install the CLI:

```
cd cli && go install ubiwhere
```

### Running

Go to the simulator folder ``cd .. && cd simulator`` and run it:

```
go run main.go
```

Alternatively, you can build and run:

```
go build main.go && ./main
```

## 🖥 CLI - Commands <a name = "commands"></a>

Run commands insde ``cli``! 

### Get the last N metrics of all variables:
```
ubiwhere read 3
```
i.e: 
```
V1: 45 | V2:65 | V3: 33 | V4: 24
V1: 44 | V2: 5 | V3: 33 | V4: 12
V1: 55 | V2:85 | V3: 98 | V4: 14
```

### Get the last N metrics of one or more variables:
```
ubiwhere read vars 2 v1 v2
```
i.e: 
```
V1: 45 | V2:65 |
V1: 44 | V2: 5 |
```

### Get the AVG of the value on one or more variables:
```
ubiwhere read avg v1 v2
```
i.e: 
```
AVG(V1): 57.3826 | AVG(V2): 73.0005 |
```

---

## 🖥 REST API <a name = "rest"></a>

### Get the last N metrics of all variables:
```
localhost:8080/ubiwhere/read/3
```
This requests gets the last 3 metrics of all variables.

### Get the last N metrics of one or more variables:
```
localhost:8080/ubiwhere/read/3/vars?var=v1&var=v2
```
This request would list the last 3 metrics of the variables V1 and V2


### Get the AVG of the value on one or more variables:
```
localhost:8080/ubiwhere/avg?var=v1&var=v2
```
This requests gets the AVG of the variables V1 and V2

NOTES: 
- Any request with query param different than ``var`` is ignored.
- Requests with var != v1, v2, v3 or v4 gets http response 400. (You need to use valid variables)

---

## Project Structure <a name = "project_structure"></a>

- **/**
    - **Simulator:**
        - **Model:** contais the data models for the Data Samples and CPU&RAM Info
        - **Controller:** contains all the functionalities/funcions of the app
        - **Rest:** router (gin) for the REST API
    - **CLI:**
        - **Model:** has the Data Samples model
        - **cmd:** the CLI funcions and configuration
    - **Database**
    
## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Since this is a challenge, database concurrency was not taken in consideration, although GORM manages this pretty well.
- I prefer a structured project because it helps understand and organize things better. 
You can check my project structure [above](#project_structure).
- References:
    - https://github.com/mackerelio/go-osstat (CPU & RAM info)