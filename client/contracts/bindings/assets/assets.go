package assets

//go:generate go-bindata -nometadata -o bind_contracts.go -pkg assets ../../truffle/contracts/consensus/token/MiningToken.sol ../../truffle/contracts/consensus/mgr/ValidatorMgr.sol ../../truffle/contracts/kns/PublicResolver.sol ../../truffle/contracts/kns/FIFSRegistrar.sol ../../truffle/contracts/kns/KNSRegistry.sol ../../truffle/contracts/oracle/ExchangeMgr.sol ../../truffle/contracts/oracle/OracleMgr.sol ../../truffle/contracts/ownership/MultiSigWallet.sol ../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxyFactory.sol ../../truffle/contracts/sysvars/SystemVars.sol ../../truffle/contracts/utils/NameHash.sol ../../truffle/contracts/utils/Strings.sol
