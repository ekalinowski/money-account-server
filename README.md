# Money Account Service

The following code is a rest api written in go to make credit and debit operations on an account.

The money-account binary file was created to run with the following command

### `./money-account`

Also to run you can use golang, or generate the docker image. The commands to develop the second option are left.

### `docker build -t money-account .`

And then

### `docker run -dp 8080:8080 money-account`

The project starts on port 8080 by default.

### Postman

Inside the miscellaneous folder there is a collection of postam to test the apis.


