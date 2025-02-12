#!/bin/bash

# gen_service.sh
SERVICE_NAME=$1

if [ -z "$SERVICE_NAME" ]; then
    echo "Usage: bash gen_service.sh <service_name>"
    exit 1
fi

echo "Generating $SERVICE_NAME client..."
cd "/d/tiktok_e-commence/gomall/rpc_gen" && cwgo client --type RPC --service "$SERVICE_NAME" --module gomall/rpc_gen --I "../idl" --idl "../idl/${SERVICE_NAME}.proto"

echo "Generating $SERVICE_NAME service..."
cd "/d/tiktok_e-commence/gomall/app/$SERVICE_NAME" && cwgo server --type RPC --service "$SERVICE_NAME" --module gomall/app/$SERVICE_NAME --pass "-use gomall/rpc_gen/kitex_gen" --I "../../idl" --idl "../../idl/${SERVICE_NAME}.proto"

echo "Generation completed."