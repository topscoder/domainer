#!/bin/bash

# Define the file to read
file="tlds.txt"

# Read the file content as a single line
content=$(cat "$file")

# Counter for the number of spaces
counter=0

# Loop through each character in the line
for (( i = 0; i < ${#content}; i++ )); do
    # Check if the character is a space
    if [[ ${content:i:1} == " " ]]; then
        # Increment the space counter
        ((counter++))

        # Print a new line if the space counter is a multiple of 10
        if ((counter % 10 == 0)); then
            echo ""
        fi
    fi

    # Print the character
    echo -n "${content:i:1}"
done

# Print a new line at the end
echo ""
