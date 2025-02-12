#! bin/bash
# gen_frontend.sh
FRONTEND_NAME=$1

if [ -z "$FRONTEND_NAME" ]; then
    echo "Usage: bash gen_frontend.sh <xxx>_page"
    exit 1
fi
echo "Generating $FRONTEND_NAME _page..."
cd "/d/tiktok_e-commence/gomall/app/frontend" && cwgo server --type HTTP --service frontend --module gomall/app/frontend -I ../../idl --idl "../../idl/frontend/${FRONTEND_NAME}_page.proto"
echo "Done"