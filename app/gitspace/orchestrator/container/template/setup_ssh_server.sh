#!/bin/sh

# Install SSH if it's not already installed
if ! command -v sshd >/dev/null 2>&1; then
    echo "OpenSSH server is not installed. Installing..."
    apt-get update
    apt-get install -y openssh-server
else
    echo "OpenSSH server is already installed."
fi

username={{ .Username }}
password={{ .Password }}
workingDir={{ .WorkingDirectory }}

# Check if the user already exists
if id "$username" >/dev/null 2>&1; then
    echo "User $username already exists."
else
    # Create a new user
    adduser --disabled-password --home "$workingDir" --gecos "" "$username"
    if [ $? -ne 0 ]; then
        echo "Failed to create user $username."
        exit 1
    fi
fi

# Set or update the user's password using chpasswd
echo "$username:$password" | chpasswd

# Configure SSH to allow this user
config_file='/etc/ssh/sshd_config'
grep -q "^AllowUsers" $config_file
if [ $? -eq 0 ]; then
    # If AllowUsers exists, add the user to it
    sed -i "/^AllowUsers/ s/$/ $username/" $config_file
else
    # Otherwise, add a new AllowUsers line
    echo "AllowUsers $username" >> $config_file
fi

# Ensure password authentication is enabled
sed -i 's/^PasswordAuthentication no/PasswordAuthentication yes/' $config_file
if ! grep -q "^PasswordAuthentication yes" $config_file; then
    echo "PasswordAuthentication yes" >> $config_file
fi

# Changing ownership of everything inside user home to the newly created user
chown -R $username .

mkdir /var/run/sshd