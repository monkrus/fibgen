# fibgen
Fibonacci sequence 

* Prerequisites : Ubuntu 20.04, Go 1.16.5, Postgre 13.3

* In your project

1. Install `go mod init github.com/monkrus/fibgen.git` 

* In your DB

1. Run `sudo passw postgres` to update password (admin)

2. Start server by running `sudo service postgresql server`

3. Update your password `sudo -u postgres psql` 
   Type `ALTER USER postgres PASSWORD ``admin;`

4. Create DB by typing `CREATE DATABASE fibo_keeper;`

5. Exit `\q`
   


