pwncheck checks your password against [pwnedpasswords.com]'s API. It does this by submitting a prefix of the hash of your password, minimizing any risk of exposing your password to the world.

Usage: `go run pwncheck.go <mypassword>`.
	
__WARNING:__ Don't keep runs of this in your shell history with your real password!
