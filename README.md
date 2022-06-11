# Chat Client

## The Chat Client Network Standard:
### Network:

- The standard can opperate over either UDP or TCP
- The standard will opperate out of default port 1234
- The standard will have 2 stages, send & confirm

### Sending A Message:
A typical request may look like

    FROM:Geof:192.168.1.10:1234/Hello how are you/Send
    
This would be a message from Geof on the ip 192.168.1.10 to a client on
the ip 192.168.1.11.

### Confirming:
To confirm the message a hash of the request will be sent back to the client by the target
    
    9b527f578f800bfc1067748a677fb35cf2e2ce9bf3ca1740774ee7be75136f0c
    
If all the hash matches on both clients it is known that a secure transmission of the message has been made.

## If sent has is correct:
Upon receiving a confirmation the client will compare the hash against a hash of the original message.
if they match it is known that the message that has been sent has been correctly received and the connection will be
terminated as no more action is needed. 

## If sent hash is incorrect:
If the confirmation hash is incorrect a Replace request can be sent:
    
    FROM:Geof:192.168.1.10:1234/Hello how are you/Replace

This will tell the receiver that the previous message wasn't received properly. A standard confirmation
will be sent back. X number of replace requests can be made before an unstable connection is determined
and the end user is notified that the message may have not been sent accurately.

## A typical CC conversation (Perspective of Benjamin)

User inputs a message:

    > Hello

Message is parsed into a CC request and a hash

    CC request = FROM:Benjamin:192.168.1.11:1234/Hello/Send
    Hash = 9b527f578f800bfc1067748a677fb35cf2e2ce9bf3ca1740774ee7be75136f0c

Sending of the CC request to Geof on 192.168.1.10:
    
    FROM:Benjamin:192.168.1.11:1234/Hello/Send

Confirmation received from Geof:

    FROM:Geof:192.168.1.10:1234/9b527f578f800bfc1067748a677fb35cf2e2ce9bf3ca1740774ee7bakedbeans/Confirm

Benjamin Checks the hash:
    
    9b527f578f800bfc1067748a677fb35cf2e2ce9bf3ca1740774ee7be75136f0c == 9b527f578f800bfc1067748a677fb35cf2e2ce9bf3ca1740774ee7bakedbeans

The hash does not match so Benjamin sends a replace request to Geof on 192.168.1.10 and forms a new hash of the message:

    FROM:Benjamin:192.168.1.11:1234/Hello/Replace

    new hash = 54e9ecdb50fdaa43bc05d963f61f82c0b58377b01e571766251fbfe4d178d6fe

Confirmation received from Geof:
    
    FROM:Geof:192.168.1.10:1234/54e9ecdb50fdaa43bc05d963f61f82c0b58377b01e571766251fbfe4d178d6fe/Confirm
    
New hash is checked:

    54e9ecdb50fdaa43bc05d963f61f82c0b58377b01e571766251fbfe4d178d6fe == 54e9ecdb50fdaa43bc05d963f61f82c0b58377b01e571766251fbfe4d178d6fe

The hashes match so the connection is closed
    


    
    