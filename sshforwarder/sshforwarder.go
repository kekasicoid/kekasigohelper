package sshforwarder

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func ForwardSSHConn(localListener net.Listener, sshConn *ssh.Client, remoteAddr string) {
	go func() {
		for {
			localConn, err := localListener.Accept()
			if err != nil {
				log.Println("Local accept error:", err)
				continue
			}
			go func() {
				remoteConn, err := sshConn.Dial("tcp", remoteAddr)
				if err != nil {
					log.Println("Remote dial error:", err)
					localConn.Close()
					return
				}
				go io.Copy(remoteConn, localConn)
				go io.Copy(localConn, remoteConn)
			}()
		}
	}()
}
