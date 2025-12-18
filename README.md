\# Go User Service



A RESTful API built with Go, Fiber, and SQLC.



\## 🔧 Tech Stack

\* \*\*Language:\*\* Go (Golang)

\* \*\*Framework:\*\* GoFiber

\* \*\*Database:\*\* MySQL + SQLC (Type-safe SQL)

\* \*\*Architecture:\*\* Layered (Handler -> Service -> Repository)



\## 🚀 How to Run



\### 1. Prerequisites

\* Go installed

\* MySQL running locally



\### 2. Database Setup

Run the following SQL in your MySQL client to create the table:

```sql

CREATE DATABASE IF NOT EXISTS userdb;

USE userdb;

CREATE TABLE users (

&nbsp;   id BIGINT AUTO\_INCREMENT PRIMARY KEY,

&nbsp;   name TEXT NOT NULL,

&nbsp;   dob DATE NOT NULL

);

