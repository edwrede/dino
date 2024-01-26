/*
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type animal struct {
	animal_ID       int
	animal_type     string
	animal_nickname string
	animal_zone     int
	animal_age      int
}

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

type DatabaseCreds struct {
	SSHHost    string // SSH Server Hostname/IP
	SSHPort    int    // SSH Port
	SSHUser    string // SSH Username
	SSHKeyFile string // SSH Key file location
	DBUser     string // DB username
	DBPass     string // DB Password
	DBHost     string // DB Hostname/IP
	DBName     string // Database name
}

func main() {
	db, sshConn, err := ConnectToDB(DatabaseCreds{
		SSHHost:    "159.65.52.97",
		SSHPort:    22,
		SSHUser:    "root",
		SSHKeyFile: "sshkeyfile.pem",
		DBUser:     "root",
		DBPass:     "LetmeIn123!",
		DBHost:     "localhost:3306",
		DBName:     "dino",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sshConn.Close()
	defer db.Close()

	if rows, err := db.Query("SELECT 1=1"); err == nil {
		for rows.Next() {
			var result string
			rows.Scan(&result)
			fmt.Printf("Result: %s\n", result)
		}
		rows.Close()
	} else {
		fmt.Printf("Failure: %s", err.Error())
	}
}

func ConnectToDB(dbCreds DatabaseCreds) (*sql.DB, *ssh.Client, error) {

	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	/*pemBytes, err := os.ReadFile(dbCreds.SSHKeyFile)
	if err != nil {
		return nil, nil, err
	}
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		return nil, nil, err
	}*/

	/*
	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User:            dbCreds.SSHUser,
		Auth:            []ssh.AuthMethod{ssh.Password("LetmeIn123!")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}

	// Connect to the SSH Server
	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", dbCreds.SSHHost, dbCreds.SSHPort), sshConfig)
	if err != nil {
		return nil, nil, err
	}

	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	mysql.RegisterDialContext("mysql+tcp", func(_ context.Context, addr string) (net.Conn, error) {
		dialer := &ViaSSHDialer{sshConn}
		return dialer.Dial(addr)
	})

	// And now we can use our new driver with the regular mysql connection string tunneled through the SSH connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", dbCreds.DBUser, dbCreds.DBPass, dbCreds.DBHost, dbCreds.DBName))
	if err != nil {
		return nil, sshConn, err
	}

	fmt.Println("Successfully connected to the db")

	return db, sshConn, err

	//connect to the database
	/*
		log.Println("attempting DB connection")
		db, err := sql.Open("mysql", "root:LetmeIn123!@tcp(159.65.52.97:3306)/dino")
		if err != nil {
			log.Fatal(err)
			log.Println("DB connection failed")
		}
		log.Println("DB connection successful")
		log.Println(db.Stats().InUse)
		defer db.Close()

		//Send query to DB
		log.Println("attempting query")
		rows, err := db.Query("SELECT * FROM dino.animals where animal_ID >= ?", 1)
		if err != nil {
			log.Fatal(err)
			log.Println("Query error")
		}
		log.Println("Query success")
		defer rows.Close()

		animals := []animal{}

		for rows.Next() {
			a := animal{}
			err := rows.Scan(&a.animal_ID, &a.animal_type, &a.animal_nickname, &a.animal_zone, &a.animal_age)
			if err != nil {
				log.Println(err)
				continue
			}
			animals = append(animals, a)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

	*/
}*/
