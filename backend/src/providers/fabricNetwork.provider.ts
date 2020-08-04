import {inject} from '@loopback/core';
import * as fabricNetwork from 'fabric-network';
import * as Client from 'fabric-client';
import * as fs from 'fs';
import {v4 as uuid} from 'uuid';
import {ContractEventListener} from 'fabric-network';
import {getLogger, Logger, } from '../utils/logger';
import {ConfigurationBindings} from '../config/binding.constants';

interface FabricNetworkProviderConfiguration {
  localMspId: string;
  mspConfigPath: string;
  commonConnectionProfilePath: string;
  network: string;
  chaincodeId: string;
}

export class FabricNetworkProvider implements fabricNetwork.Contract {
  addDiscoveryInterest(interest: fabricNetwork.DiscoveryInterest): fabricNetwork.Contract {
    throw new Error("Method not implemented.");
  }
  private logger: Logger = getLogger(FabricNetworkProvider.constructor.name);
  wallet: fabricNetwork.Wallet;
  gateway: fabricNetwork.Gateway;
  network: fabricNetwork.Network;
  contract: fabricNetwork.Contract;

  listeners: {[key: string]: ContractEventListener} = {};

  constructor(
    @inject(ConfigurationBindings.FABRIC_NETWORK_CONFIGURATION)
    private fabricNetworkConfiguration: FabricNetworkProviderConfiguration,
  ) {}

  async init() {
    await this.initWallet();
    await this.initNetworkConnection();
    await this.initContract();
  }

  async deinit() {
    this.gateway.disconnect();
  }

  async initWallet() {
    this.logger.debug('initWallet');
    this.logger.debug("tracing mspConfigPath", this.fabricNetworkConfiguration.mspConfigPath);
    this.logger.debug("localMspId", this.fabricNetworkConfiguration.localMspId);
    const identity = fabricNetwork.X509WalletMixin.createIdentity(
      this.fabricNetworkConfiguration.localMspId,
      this.getCertificateFromMSP(
        this.fabricNetworkConfiguration.mspConfigPath,
        0,
      ),
      this.getPrivateKeyFromMSP(
        this.fabricNetworkConfiguration.mspConfigPath,
        0,
      ),
    );
    this.wallet = new fabricNetwork.InMemoryWallet();
    await this.wallet.import(
      this.fabricNetworkConfiguration.localMspId,
      identity,
    );
  }

  async initNetworkConnection() {
    this.logger.debug('initNetworkConnection');
    const gatewayOptions: fabricNetwork.GatewayOptions = {
      identity: this.fabricNetworkConfiguration.localMspId,
      wallet: this.wallet,
    };
    this.gateway = new fabricNetwork.Gateway();
    await this.gateway.connect(
      this.fabricNetworkConfiguration.commonConnectionProfilePath,
      gatewayOptions,
      );
  }

  async initContract() {
    this.logger.debug('initContract');
    this.network = await this.gateway.getNetwork(
      this.fabricNetworkConfiguration.network,
      );
      this.logger.debug("localMspIdddd2");
    this.contract = this.network.getContract(
      this.fabricNetworkConfiguration.chaincodeId,
    );
  }

  getCertificateFromMSP(mspConfigPath: string, index: Number): string {
    const files = fs.readdirSync(`${mspConfigPath}/signcerts`);
    let currentIndex = 0;
    for (const filename in files) {
      if (files[filename].endsWith('.pem')) {
        const certPath = `${mspConfigPath}/signcerts/${files[filename]}`;
        if (index === currentIndex) {
          this.logger.trace(`Found this certPath: ${certPath}`);
          return fs.readFileSync(certPath).toString();
        }
        currentIndex++;
      }
    }
    throw new Error(`Could not find any certPath at index ${index}`);
  }

  getPrivateKeyFromMSP(mspConfigPath: string, index: Number): string {
    const files = fs.readdirSync(`${mspConfigPath}/keystore`);
    let currentIndex = 0;
    for (const filename in files) {
      if (files[filename].endsWith('_sk')) {
        const privateKeyPath = `${mspConfigPath}/keystore/${files[filename]}`;
        if (index === currentIndex) {
          this.logger.trace(`Found this privateKeyPath: ${privateKeyPath}`);
          return fs.readFileSync(privateKeyPath).toString();
        }
        currentIndex++;
      }
    }
    throw new Error(`Could not find any privateKey at index ${index}`);
  }

  // Contract interface
  createTransaction(name: string): fabricNetwork.Transaction {
    const transaction = this.contract.createTransaction(name)
    transaction.setEndorsingOrganizations(this.fabricNetworkConfiguration.localMspId)
    return transaction;
  }
  evaluateTransaction(name: string, ...args: string[]): Promise<Buffer> {
    return this.contract.evaluateTransaction(name, ...args);
  }
  submitTransaction(name: string, ...args: string[]): Promise<Buffer> {
    return this.contract.submitTransaction(name, ...args);
  }

  addContractListener(
    listenerName: string,
    eventName: string,
    callback: (
      error: Error,
      event?: Client.ChaincodeEvent | Client.ChaincodeEvent[],
      blockNumber?: string,
      transactionId?: string,
      status?: string,
    ) => Promise<void> | void,
    options?: fabricNetwork.EventListenerOptions,
  ): Promise<fabricNetwork.ContractEventListener> {
    return this.contract.addContractListener(listenerName, eventName, callback);
  }

  async addEventListener(
    eventName: string,
    onEvent: (
      error: Error,
      event?: Client.ChaincodeEvent | Client.ChaincodeEvent[],
      blockNumber?: string,
      transactionId?: string,
      status?: string,
    ) => Promise<void> | void,
  ): Promise<string> {
    const listenerUuid = uuid();
    const contractEventListenerHandle = await this.contract.addContractListener(
      listenerUuid,
      eventName,
      onEvent,
    );
    this.listeners[listenerUuid] = contractEventListenerHandle;
    return Promise.resolve(listenerUuid);
  }

  async removeEventListener(listenerUuid: string) {
    this.listeners[listenerUuid].unregister();
    delete this.listeners[listenerUuid];
  }
}
