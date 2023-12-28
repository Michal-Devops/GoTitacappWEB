#!/bin/bash

# Definiowanie zmiennych
INSTANCE_TYPE="t2.micro"
AMI_ID="ami-0ff8a91507f77f867"  # Zastąp własnym AMI
KEY_NAME="MyEC2KeyPair"    # klucze
SECURITY_GROUP="my-sec-group" #  moja grupa zabezpieczen 
SUBNET_ID="subnet-xxxxxxxx"   # tu normalnie jest moje id 
# Utworzenie instancji EC2
INSTANCE_ID=$(aws ec2 run-instances \
    --image-id $AMI_ID \
    --instance-type $INSTANCE_TYPE \
    --key-name $KEY_NAME \
    --security-group-ids $SECURITY_GROUP \
    --subnet-id $SUBNET_ID \
    --query 'Instances[0].InstanceId' \
    --output text)

echo "Instancja EC2 utworzona: $INSTANCE_ID"

# Oczekiwanie na uruchomienie instancji
aws ec2 wait instance-running --instance-ids $INSTANCE_ID
echo "Instancja EC2 jest teraz uruchomiona."

# Pobranie adresu IP instancji
IP_ADDRESS=$(aws ec2 describe-instances \
    --instance-ids $INSTANCE_ID \
    --query 'Reservations[0].Instances[0].PublicIpAddress' \
    --output text)

echo "Adres IP instancji EC2: $IP_ADDRESS"


