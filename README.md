This Go script ssh's into a remote machine and runs the command user enters. The program has been written in GoLang and uses ssh package to do the aforementioned task.

The code comprises of two validation functions, one ssh function and a main function. The validation functions validate the format of the IP address as well as the format of the overall address. The ssh function accesses the remote machine runs the command.

NOTE: This can also be achieved by crypto/ssh package provided by go out of the box.
