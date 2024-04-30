# Root directory
RECEIPT_PROVER_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

cd $RECEIPT_PROVER_PATH
go build -o build/receipt_prover main/main.go
