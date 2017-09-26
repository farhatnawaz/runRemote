This program accesses a machine remotely and runs the command user enters. The program has been written in Go language and uses ssh package to do the aforementioned task.

The code comprises of two validation functions, one ssh function and a main function. Two validation functions are used to validate the format of the IP address as well as the format of the overall address.

The ssh function is the actual one that is used to access the machine remotely using the address provided by the user and then run the command.

NOTE: This can also be achieved by crypto/ssh package provided by go out of the box.
