# Go for DevOps: A Beginner's Guide

This repository is a collection of Go (Golang) scripts and examples tailored for DevOps practices. It includes basic Go programs, configuration management, and Kubernetes interaction scripts. Below is an overview of the files and commands to help you understand how Go can be used in a DevOps context.

---

## 🛠️ Install Go (Golang)  

Before running this project, you need to install **Go**. Follow the steps below based on your operating system:  

### **🔹 Windows**  
1. Download the latest **Go installer** from [golang.org](https://go.dev/dl/).  
2. Run the `.msi` installer and follow the on-screen instructions.  
3. Verify the installation by running:  
   ```sh
go version
   ```


### **🔹 Macos**  

```sh
brew install go

```

### **🔹 Linux**  

Use snap to install the latest version dont go with wget tar.gz
```sh
sudo snap install go

```

🤝 Contributing

Feel free to submit PRs or open issues for enhancements!


### **Why This is a Good README?**  
✅ **Covers Project Setup** (`go mod init`, `go mod tidy`)  
✅ **Explains Each File** in the directory  
✅ **Step-by-Step Guide** on running Go scripts  
✅ **Keeps It Clean & Copy-Paste Ready*

## Files in This Directory

- **config.yaml**: A configuration file used for managing settings in Go scripts.
- **go.mod**: Defines the Go module, including the module path and dependencies.
- **go.sum**: Contains cryptographic checksums for verifying dependency integrity.
- **helloworld.go**: A simple "Hello, World!" program to get started with Go.
- **learn.txt**: Notes and learning points from my Go and DevOps journey.
- **readapi.go**: A script demonstrating how to read data from an API.
- **readconfigfile.go**: A script that reads and parses the `config.yaml` file.
- **showk8sresrc.go**: A script that interacts with Kubernetes resources (e.g., listing or managing them).
- **txt**: A directory containing additional text files or resources.

---

## Essential Go Commands for DevOps

Here are some key Go commands and their purposes, especially useful for DevOps workflows:

| **Command**               | **Purpose**                                                                 |
|---------------------------|-----------------------------------------------------------------------------|
| `go mod init <name>`       | Initializes a Go module (creates `go.mod`).                                 |
| `go.mod`                   | Stores project details and dependencies.                                    |
| `go get <package>`         | Downloads a dependency and updates `go.mod`.                                |
| `go.sum`                   | Security file that verifies dependency integrity.                           |
| `go mod tidy`              | Removes unused dependencies and fixes missing ones.                         |
| `go run <file.go>`         | Runs the Go program.                                                        |
| `go build <file.go>`       | Compiles the Go program into an executable binary.                          |
| `go test`                  | Runs tests in the current directory.                                        |
| `go fmt <file.go>`         | Formats the Go code to adhere to standard conventions.                      |

---

## How to Use This Repository

Clone the project

```bash
  git clone https://github.com/mehboobpatel/Go-Lang-ForDevOps
```

Go to the project directory

```bash
  cd Go-Lang-ForDevOps
```

Initialize your WorkDir

```bash
go mod init <yourprojectname>
```
- This will generate go.mod

To Download any Additional package 
```bash
go get k8s.io/client-go@latest
```
- This will generate go.sum

If your code imports a package that is not in go.mod, it adds it and updates go.sum.
If your go.mod file contains imports that are no longer used, it removes them from go.mod and go.sum
```bash
go mod tidy
```

To print a helloworld with your name as input  
```bash
go run helloworld.go Maheboob
```

To compile into an executable binary:

```bash
go build helloworld.go 
./helloworld Maheboob
```

To fetch & read data from an API and processes it.
```bash
go run readapi.go
```


To read a file.
```bash
go run readconfigfile.go
```

To Interact with Kubernetes & List with resources
```bash
go run showk8sresrc.go
```

🤝 Contributing

Feel free to submit PRs or open issues for enhancements!


### **Why This is a Good README?**  
✅ **Covers Project Setup** (`go mod init`, `go mod tidy`)  
✅ **Explains Each File** in the directory  
✅ **Step-by-Step Guide** on running Go scripts  
✅ **Keeps It Clean & Copy-Paste Ready*