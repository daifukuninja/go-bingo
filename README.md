# go-bingo

## Outlines

- go + cli + BINGO
- generate number list 1 to 75
- generate 5 sets number groups (15 numbers per group)
- display groups and numbers shapes 3 x 5 (B-I-N-G-O)
- change group color in order
- after group is decided, change number color in order
- Finally, pick a number

## Features

## How to debug

definition to `launch.json`

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Attach to Bubbletea",
            "type": "go",
            "debugAdapter": "dlv-dap",
            "request": "attach",
            "mode": "remote",
            "remotePath": "${workspaceFolder}",
            "port": 2345,
            "host": "127.0.0.1",
        },
    ]
}
```

`dlv debug --headless --listen=:2345 .`

Run `Attach to Bubbletea` on Run and Debug tab.
