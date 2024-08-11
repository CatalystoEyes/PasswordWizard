# GoPassword

Welcome to the GoPassword application! This service, built with Go and Vue, generates random passwords using bcrypt to enhance security.

## Overview

GoPassword provides a simple interface to create strong, random passwords that are hashed with bcrypt. This ensures that your passwords are not only unique but also securely stored.

## Features

- Generates random passwords of customizable length
- Passwords are hashed using bcrypt
- Simple and intuitive user interface built with Vue.js

## Installation

Follow the steps below to set up GoPassword on your local machine.

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Node.js](https://nodejs.org/) (recommended version 14 or higher)
- [Vue CLI](https://cli.vuejs.org/)

### Steps to Install

1. **Clone the Repository**

   Open your terminal and run:

   ```bash
   git clone https://github.com/CatalystoEyes/GoPassword.git
   ```

2. **Navigate to the Project Directory:**
   `cd GoPassword`
3. **Build the Docker Image:**

```bash
  docker build -t GoPassword .
```

4. **Run the Docker Container:**

```bash
docker run -p 8080:8080 GoPassword
```

### Finish!

Now you can access the GoPassword application at http://localhost:8080 in your web browser!
