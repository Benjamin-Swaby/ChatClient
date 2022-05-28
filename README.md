# Chat Client

## The Chat Client Network Standard:
### Network:

- The standard can opperate over either UDP or TCP
- The standard will opperate out of default port 1234
- The standard will have 2 stages, send & confirm

### Sending A Message:
    A typical request may look like

    `FROM:Geof:192.168.1.10/Hello how are you/192.168.1.11/Send`
    
    This would be a message from Geof on the ip 192.168.1.10 to a client on
    the ip 192.168.1.11.

### Confirming:
    Confirmation will be a seperate request consisting of their local ip
    the senders ip and a hash of the original message. i.e:
    
    `FROM:Benjamin:192.168.1.11/c954abb0fda4c/192.168.1.10/Confirm`
    
    If all the details match on both clients it is known that a secure transmission of the message has been made.

    
    