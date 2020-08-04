export * from 'child_process';
export const configurationGeneral = {
  sqlstore: {
    datasource: {
      username: process.env.POSTGRES_USER || 'postgres',
      password: process.env.POSTGRES_PASSWORD || 'postgres',
      connector: 'postgresql',
      host: process.env.POSTGRES_HOST || 'localhost',
      port: 5432,
      database: 'blockchainproject',
    },
  },
  logger: {
    level: 'debug',
    pattern: '[%d] [%p] %c - %m',
  },

  fabricNetworkConfiguration: {
    localMspId: 'partya',
    mspConfigPath: `/Users/admin/go/src/test/Blockchain-Project/network/crypto-config/peerOrganizations/partya.example.com/users/Admin@partya.example.com/msp`,
    /* mspConfigPath: `/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/partya.example.com/users/Admin@partya.example.com/msp`, */
    commonConnectionProfilePath: '/Users/admin/go/src/test/Blockchain-Project/network/network-config.yaml',
    network: 'irs',
    chaincodeId: 'irscc',
  },
};
export default configurationGeneral;