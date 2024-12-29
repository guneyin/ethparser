
# Ethereum Transaction Parser
[![Go Coverage](https://github.com/guneyin/ethparser/wiki/coverage.svg)](https://raw.githack.com/wiki/guneyin/ethparser/coverage.html)

Parse Ethereum transactions by address

## Run Locally

Clone the project

```bash
  git clone https://github.com/guneyin/ethparser.git
```

Go to the project directory

```bash
  cd ethparser
```

Init project

```bash
  make init
```

Start the server

```bash
  ./ethparser --port=3000
```


## API Reference

### Get current block

Get latest block from ethereum network

```http
  GET /current-block
```

##### Response
```json
{
    "result": 21506906
}
```

### Subscribe

```http
  GET /subscribe?addr=${addr}
```

| Parameter | Type     | Description                   |
|:----------|:---------|:------------------------------|
| `addr`    | `string` | **Required**. Address of user |

##### Response
```json
{
    "result": true
}
```

### Transactions

```http
  GET /transactions?addr=${addr}
```

| Parameter | Type     | Description                   |
|:----------|:---------|:------------------------------|
| `addr`    | `string` | **Required**. Address of user |

##### Response
```json
{
    "result": [
        {
            "blockHash": "0xfa8bdc89435a4890ea28d184d28776c8d7c27f23b2cf3483d0f8ac28207b9bcb",
            "blockNumber": "0x1481c79",
            "from": "0xe97881a663a97b1fecbfe51a6318f2d53c7a49cd",
            "gas": "0x2760d",
            "gasPrice": "0x35182b2e4",
            "hash": "0x5fc7bd7f1c1332ebb56c569823a9e64fb51f9ce0262cb50d1c535bab0f7d9415",
            "input": "0x122067ed000000000000000000000000000000000000000000000000000000000000000000000000000000000000000059c38b6775ded821f010dbd30ecabdcf84e0475600000000000000000000000000000000000000000000002418750b61872c000000000000000000000000000000000000000000000000000169bb965299bde0000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000006770537b00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000141f9840a85d5af5bf1d1762f925bdaddc4201f984000000000000000000000000",
            "nonce": "0x31ba",
            "to": "0x51c72848c68a965f66fa7a88855f9f7784502a7f",
            "transactionIndex": "0x0",
            "value": "0x0",
            "type": "0x2",
            "chainId": "0x1",
            "v": "0x0",
            "r": "0x641bfd06c0d1daee51650ce4c9a1c4d1196e224cf2cf4e63dfcabbfbd009f353",
            "s": "0x5d8725f5d2455cdb3d9bbad747184da96d62da9a20f4fa5d29992f4d2ff931be"
        }
    ]
}
```

