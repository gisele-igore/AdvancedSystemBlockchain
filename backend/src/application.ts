import {BootMixin} from '@loopback/boot';
import {ApplicationConfig} from '@loopback/core';
import {RepositoryMixin} from '@loopback/repository';
import {RestApplication} from '@loopback/rest';
import {RestExplorerComponent, RestExplorerBindings} from '@loopback/rest-explorer';
import {ServiceMixin} from '@loopback/service-proxy';
import * as hfc from 'fabric-client';

import {shutdownLogger} from './utils/logger';
import {configurationGeneral} from './config';
import {ProviderBindings} from './providers';
import {
  
  CompagnieAssuranceService,
  ServiceBindings,
  
} from './services';

export class BlockchainProjectBackendApplication extends BootMixin(
  ServiceMixin(RepositoryMixin(RestApplication)),
) {
  constructor(options: ApplicationConfig = {}) {
    super(Object.assign(options, configurationGeneral));

    this.projectRoot = __dirname;
    // Customize @loopback/boot Booter Conventions here
    this.bootOptions = {
      controllers: {
        // Customize ControllerBooter Conventions here
        dirs: ['controllers'],
        extensions: ['.controller.js'],
        nested: true,
      },
    };
    this.basePath('/api');
    this.component(RestExplorerComponent);

    this.bind('datasources.config.sqlstore').to(
      configurationGeneral.sqlstore.datasource,
    );
    this.bind(ProviderBindings.SOCKET_IO_SERVER).to(this.socketIOServer);
    this.bind(ProviderBindings.HYPER_LEDGER_FABRIC).to(hfc);

    this.bind(ServiceBindings.COMPAGNIE_ASSURANCE_SERVICE).toClass(
      CompagnieAssuranceService,
    );
    /* const fse = require('fs-extra')
    async function copyCryptoFiles(){
      try {
        await fse.copySync('/etc/hyperledger/fabric','./cryptogen')
        console.log("Success")
      } catch (error) {
        console.error(error)
      }
    }

    if (process.env.TELEPRESENCE_CONTAINER_NAMESPACE === undefined){
      copyCryptoFiles()
    } */
  }

  // never called for now
  async stop(): Promise<void> {
    super.stop().then(
      () => shutdownLogger(),
      () => shutdownLogger(),
    );
  }
}
