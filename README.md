# Cardano Native Asset Transfer

This is an experimental repo that aims to construct a valid unsigned Cardano transaction body that transfers Cardano Native Asset tokens.

## Token Asset ID

Genius Yield Token is a native token on Cardano.

The asset ID is a pair of the policy ID and the asset name.

- Policy Id: `dda5fdb1002f7389b33e036b6afee82a8189becb6cba852e8b79b4fb`
- Asset Name: `(333) GENS (0014df1047454e53)`
- [Explorer record for this address][token on explorer]
- [Adastat Record][adastat record]

[token on explorer]: https://cardanoscan.io/token/dda5fdb1002f7389b33e036b6afee82a8189becb6cba852e8b79b4fb0014df1047454e53?address=addr1vx4us3sfrqdfswtyaf48w0vl8n04wvyf3y4gzy63y9y2racekldk3
[adastat record]: https://adastat.net/tokens/dda5fdb1002f7389b33e036b6afee82a8189becb6cba852e8b79b4fb0014df1047454e53

Address controlling these assets: `addr1vx4us3sfrqdfswtyaf48w0vl8n04wvyf3y4gzy63y9y2racekldk3`

The unspent UTXOs for this address are in the [project config](./cmd/build-tx/config.yml)
