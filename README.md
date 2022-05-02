# react-microservice

## Description

This is a microservice example for the react-app.

## Usage

1. Run database server:

```bash
make run_db
```

2. Create posts database:

```bash
make createdb_posts
```

3. Create comments database:

```bash
make createdb_comments
```

4. Install frontend dependencies

```bash
cd frontend
yarn
```

5. Run posts microservice

```bash
make run_posts_microservice
```

6. Run comments microservice

```bash
make run_comments_microservice
```

7. Run frontend

```bash
make run_frontend
```
