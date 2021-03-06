# API-Go-React

[![LinkedIn][linkedin-shield]][linkedin-url]


[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/emmanuel-palacio/

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
          <ul>
            <li><a href="#back-end">Back-End</a></li>
            <ul>
              <li><a href="#mysql-and-docker">Mysql and Docker</a></li>
              <li><a href="#go-and-its-dependencies">Go and its dependencies</a></li>
            </ul>
            <li><a href="#front-end">Front-End</a></li>
            <ul>
              <li><a href="#mysql-and-docker">Mysql and Docker</a></li>
              <li><a href="#go-and-its-dependencies">Go and its dependencies</a></li>
            </ul>
          </ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

Creation of a simple web application to record the entry and exit of your employees as well as manage their information.

### Built With
The back-end part is built with the following technologies

* [Go](https://golang.org/)
* [Mysql](https://www.mysql.com/)
* [Docker](https://www.docker.com/)

The front-end part is built with the following technologies.

* [Javascript](https://golang.org/)
* [React](https://www.mysql.com/)
* [HTML](https://lenguajehtml.com/)
* [CSS3](https://lenguajecss.com/)

## Getting Started

### Prerequisites

### Back-End
#### Mysql and Docker

* You have Docker installed, if you possibly don't have it installed, you can install it for [Ubuntu18.08](https://www.digitalocean.com/community/tutorials/como-instalar-y-usar-docker-en-ubuntu-18-04-1-es) or [Ubutu20.04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04-es) if by any chance you have another OS or another version of ubuntu, you can check the installation on the official [docker](https://www.docker.com/) page.

To start a MySQL database with docker the process couldn't be easier. Just take the template of the Docker Compose configuration file that we leave below and invoke the command that will carry out the entire process.
```
version: '2'
services:
  mysql:
    image: 'mysql: latest'
    restart: always
    volumes:
      - './mysql_data:/var/lib/mysql'
    environment:
      - MYSQL_ROOT_PASSWORD = secure_pass_here
    ports:
      - '3306: 3306'
```

Save the file locally with the name ```docker-compose.yml```  in a directory called ```<user_home>/docker/mysql``` if you are on Unix, or ```<user_home>\docker\mysql``` if you are on Windows, if there is no such route, just create it. The configuration saves the data in the ```./mysql_data``` subdirectory permanently, otherwise the hosted databases will be lost on each reboot of the container.

Then, locate yourself with a command terminal in the same directory where you save the ```docker-compose.yml``` file and run:

```docker-compose up -d```

The system will first download the image from the?? public Docker repositories ??when it sees that it is not available locally and then it will proceed to boot and open the port indicated in the configuration file. The ```-d``` option?? ??tells Docker to run the service in the background.  From then on, use your preferred graphical interface to connect to the database.

To verify that the service is up and the container is running normally, try running the following docker command that shows the containers running:

* docker ps

| CONTAINER ID   | IMAGE        |  COMMAND               | CREATED      |  STATUS        |  PORTS                                                 | NAMES        |
| :---           |     :---:    |          -----         |:---          |     :---:      |  ----------------------------------------------------  |:---          |
| 46f4b96d6b2d   | mysql:latest | "docker-entrypoint.s???" |  4 days ago  | Up 3 hours     | 0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   | mysql_mysql_1|


In my case, I am on a ubuntu 18.04 operating system and to run mysql and directly mount a docker terminal, I run the following command

```docker exec -it mysql_mysql_1 bash```

#### Go and its dependencies

1. Install third-party modules first.
```
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
```

2. For security reasons, we cannot share the credentials of the database where they are stored in a file ``` .env``` inside the Back-End directory. which is not in this repository, the file contains the following environment variables.
```
user="UserMysql"
pass="password"
host="hostlocal"
port="Number of the port"
db_name="Name of the data base"
```

### Front-End
To install the Front-End dependencies we simply execute the following command
* npm
  ```sh
  npm install
  ```
  
## Installation
### Back-End
1. Compile all the .go files and you can do it just by typing this command.
```
bo build
```
Generates a ```Back-End``` file that is the executable of the previously compiled files.

2. Then run the executable file.
```
./Back-End
```
### Front-End
After installing the dependencies, proceed to run the program
```
./npm start
```
  
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

  
## Contact

Your Name - [@SixTanDeveloper](https://twitter.com/SixTanDeveloper) - sixtandev@gmail.com

Project Link: [API Go and React](https://github.com/SixTanDev/API-Go-React)
