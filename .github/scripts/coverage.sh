#!/bin/bash

# Define color codes for terminal output
HEADER='\033[95m'
OKBLUE='\033[94m'
OKCYAN='\033[96m'
OKGREEN='\033[92m'
WARNING='\033[93m'
FAIL='\033[91m'
ENDC='\033[0m'
BOLD='\033[1m'
UNDERLINE='\033[4m'

# Check if the correct number of arguments is provided
if [ "$#" -lt 2 ]; then
    echo "Usage: $0 [threshold] [go-coverage-report]"
    exit 1
fi

# Assign threshold and report file from command-line arguments
threshold=$1
report=$2

# Run the Go coverage tool and capture the output
output=$(go tool cover -func "$report")

# Extract the coverage percentage from the output
percent_coverage=$(echo "$output" | tail -n 1 | awk '{print $3}' | sed 's/%//')

# Print the coverage percentage
echo -e "${BOLD}Coverage: ${percent_coverage}%${ENDC}"

# Check if the coverage is below the threshold
if (( $(echo "$percent_coverage < $threshold" | bc -l) )); then
    echo -e "${BOLD}${FAIL}Coverage below threshold of ${threshold}%${ENDC}"
    exit 1
fi
