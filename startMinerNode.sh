#!/bin/bash


geth --networkid 1444 --datadir "./node1/data" --bootnodes enode://a4f1d222c6f8d9cc6c3dfab0184e657f03846383e74f744d44003ba1122d60aa492eb67f897cc80185c96aee0b43464b4fe67b1db4cfdc780dee693a20e0ad38@127.0.0.1:0?discport=30301 --port 30303 --ipcdisable --syncmode full --http --allow-insecure-unlock --http.corsdomain "*"  --authrpc.port 8547 --unlock 0x58f1013870829B2d91d91b50F5ea3A8A6B4a4683 --password "./node1/password.txt" -mine console

