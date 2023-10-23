package contracts

const SidraTokenAbiString = `
[
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_count",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "ActiveMiners",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "address",
        "name": "_wallet",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "Converted",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "ConvertedSupply",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "address",
        "name": "_miner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "Mined",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "address",
        "name": "_wallet",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "bool",
        "name": "_status",
        "type": "bool"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "MinerStatus",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "address",
        "name": "_to",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "MintedByOwner",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "Paused",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "TokenSupply",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "address",
        "name": "_from",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "_to",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "_value",
        "type": "uint256"
      }
    ],
    "name": "Transfer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs":
    [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "_at",
        "type": "uint256"
      }
    ],
    "name": "Unpaused",
    "type": "event"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "_addr",
        "type": "address"
      }
    ],
    "name": "addMiner",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "_owner",
        "type": "address"
      }
    ],
    "name": "balanceOf",
    "outputs":
    [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "balances",
    "outputs":
    [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address[]",
        "name": "_addrs",
        "type": "address[]"
      }
    ],
    "name": "batchAddMiner",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address[]",
        "name": "_addrs",
        "type": "address[]"
      }
    ],
    "name": "batchRemoveMiner",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      }
    ],
    "name": "convert",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "convertedSupply",
    "outputs":
    [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "decimals",
    "outputs":
    [
      {
        "internalType": "uint8",
        "name": "",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "lastMiningTime",
    "outputs":
    [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "mine",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "miner",
    "outputs":
    [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "_to",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      }
    ],
    "name": "mint",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "name",
    "outputs":
    [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs":
    [
      {
        "internalType": "contract Owner",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "pause",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "paused",
    "outputs":
    [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs":
    [
      {
        "internalType": "address",
        "name": "_addr",
        "type": "address"
      }
    ],
    "name": "removeMiner",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "symbol",
    "outputs":
    [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalMiners",
    "outputs":
    [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalSupply",
    "outputs":
    [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "unpause",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "wac",
    "outputs":
    [
      {
        "internalType": "contract WalletAccessControl",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]
`
