# pw
Password manager CLI tool for macos

## obligatory disclaimer
I don't recommend anyone actually use this pw manager CLI I wrote (I don't).
This was created purely for fun as a learning exercise and not intended to be
anything more. While I have done my best to take security in to consideration,
I have very little experience in security. There are plenty of pw managers out
there, use those instead!

## usage
Storing a new password:
```
// pw new service password
-> pw new youtube password123
new password saved for youtube
```

Retrieving a password (copies to system clipboard):
```
// pw get service
-> pw get youtube
password copied to clipboard!
```


Editing a password:
```
// pw edit service newPassword
-> pw edit youtube myNewPassword
youtube password edited
```

Deleting a password:
```
// pw delete service
-> pw delete youtube
password for youtube deleted
-> pw get youtube
service doesn't exist
```

## how it works
Service and password data is encrypted and stored locally in json in the user's
`~/.data/pw` directory. Each time the user adds or requests data, the values
are encrypted/decrypted before or after the transaction. The data is managed in
a hash map and subsequently written to json. Requesting a password for a
service results in it being copied to clipboard, via apple's pbcopy utility.

The encryption key is randomly generated and saved in apple's keychain access
app, which is only accessible by password. This ensures that the password data
can't be decrypted if an attacker gains gets a copy of the filesystem. It does
not prevent access to the password data if the attacker knows the user's
password.
