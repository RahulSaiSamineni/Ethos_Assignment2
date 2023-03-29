# CS 587 Assignment 2




- Name: Rahul Sai Samineni

- NetID: rsamin4@uic.edu

- UIN: 652890317





## Instructions


- Open Virtual Box and add import Fedora file and configure the network adapters to receive ping.
- Download putty and SSH and login using credentials. so that you can run go files.
- Open a new terminal window and clone the files.

 

- Run the commands below to start the RPC server

```console

sudo -E ethosKillAll

rm -r ethos myRpc myRpcIndex myRpc.go

make && make install

cd client

sudo -E ethosRun

```

- Open a second terminal window and run the commands below to start the RPC client

```console

cd client

etAl client.ethos

bankingclient


```

- Open a third terminal window and run the commands below to check the logs

```console

cd client

ethosLog .

```




- Use the below command to generate a .txt file for logs(Provided different log files for different scenerios)

```console

cd client

ethosLog . >> logfile.txt

```




