# react-go-tutorial

A simple CRUD todo list app using
- Go for backend
- React for frontend
  - [ChakraUI](https://v2.chakra-ui.com/) for component library
  - [TanStack Query](https://tanstack.com/query/latest) for data-fetching library
- free MongoDB Atlas [cluster](https://cloud.mongodb.com/v2/66d3cccbe605cd7628bda426#/clusters/detail/Cluster0) for database
- [Railway](https://railway.app/project/b149945c-ceb8-4975-9a63-818712672054) for deployment and hosting

Tutorial author: <https://github.com/burakorkmez/react-go-tutorial>

## setup

1. setup Go

    ```bash
    go install github.com/air-verse/air@latest
    go get github.com/gofiber/fiber/v2
    go get github.com/joho/godotenv
    go get go.mongodb.org/mongo-driver/mongo

2. setup React

    ```bash
    cd ./client
    npm install
    ```

## usage

1. create an [`.env` file](.env.sample)

2. start the Go app using `air` (live reload for Go apps)

    ```bash
    ❯ air
    
      __    _   ___  
     / /\  | | | |_) 
    /_/--\ |_| |_| \_ v1.52.3, built with Go go1.23.0
    
    watching .
    !exclude tmp
    building...
    running...
    hello world
    Connected to MongoDB Atlas
    
    ┌───────────────────────────────────────────────────┐ 
    │                   Fiber v2.52.5                   │ 
    │               http://127.0.0.1:4000               │ 
    │       (bound on host 0.0.0.0 and port 4000)       │ 
    │                                                   │ 
    │ Handlers ............. 5  Processes ........... 1 │ 
    │ Prefork ....... Disabled  PID ............. 44620 │ 
    └───────────────────────────────────────────────────┘ 
    ```

3. start the React app

    ```bash
    cd ./client
    npm run dev
    ```

4. when production is ready, build React app and push to repo

    ```bash
    cd ./client
    npm run build
    # push to git
    ```
