# PasswordWizard

Welcome to the PasswordWizard application! This service, built with Golang and Vue, generates random passwords using bcrypt to enhance security.

## Overview

PasswordWizard provides a simple interface to create strong, random passwords that are hashed with bcrypt. This ensures that your passwords are not only unique but also securely stored.

## Features

- Generates random passwords of customizable length
- Passwords are hashed using bcrypt
- Simple and intuitive user interface built with Vue.js

## Installation

Follow the steps below to set up PasswordWizard on your local machine.

### Prerequisites

Make sure you have the following installed:

- [Node.js](https://nodejs.org/) (recommended version 14 or higher)
- [Vue CLI](https://cli.vuejs.org/)
- [Golang]

### Steps to Install

1. **Clone the Repository**

   Open your terminal and run:

   ```bash
   git clone https://github.com/CatalystoEyes/PasswordWizard.git
   ```

2. **Navigate to the Project Directory:**
   `cd PasswordWizard`
3. **Build the Docker Image:**

```bash
  docker build -t PasswordWizard .
```

4. **Run the Docker Container:**

```bash
docker run -p 8080:8080 PasswordWizard
```
