import {BlockchainProjectBackendApplication} from './application';
import {ApplicationConfig} from '@loopback/core';
import {configureLogger, getLogger} from './utils/logger';
import * as socket from 'socket.io';
import {FabricNetworkProvider, DatabaseEventRegister} from './providers';
import {BuilderBindings} from './global.bindings';
import {ProviderBindings} from './providers';
import {MySequence} from './sequence';
import {configurationGeneral} from './config/configurationGeneral';
import {ConfigurationBindings} from './config/binding.constants';
import { WebSocketListener } from './providers/webSocketListener.provider';

export {BlockchainProjectBackendApplication};

const socketIOPort = 7000;

export async function main(options: ApplicationConfig = {}) {
  configureLogger();
  const app = new BlockchainProjectBackendApplication(options);
  const logger = getLogger();
  const socketIOServer = socket();
  socketIOServer.listen(socketIOPort);

  logger.info(`Socket IO : Listening on port : ${socketIOPort}`);
  logger.info(`Socket IO : Watching path : ${socketIOServer.path()}`);

  await app.boot();
  /* await app.migrateSchema();
  logger.info(`Socket IO : Watching path : ${socketIOServer.path()}`);
 */

  app.bind(ProviderBindings.SOCKET_IO_SERVER).to(socketIOServer);
  app
    .bind(BuilderBindings.FABRIC_NETWORK_BUILDER)
    .toClass(FabricNetworkProvider);
  app
    .bind(ConfigurationBindings.FABRIC_NETWORK_CONFIGURATION)
    .to(configurationGeneral.fabricNetworkConfiguration);
    logger.info(`Socket IO : Watching path : 3`);
  const fabricNetworkProvider = app.getSync<FabricNetworkProvider>(
    BuilderBindings.FABRIC_NETWORK_BUILDER,
  );
  await fabricNetworkProvider.init();
  app.bind(ProviderBindings.FABRIC_NETWORK).to(fabricNetworkProvider);
  app.bind(ProviderBindings.DATABASE_EVENT_REGISTER).toClass(DatabaseEventRegister);
  
  app.bind(ProviderBindings.WEB_SOCKET_LISTENER).toClass(WebSocketListener);
  
  const webSocketListener = app.getSync<WebSocketListener>(
    ProviderBindings.WEB_SOCKET_LISTENER,
  );
  const databaseEventRegister = app.getSync<DatabaseEventRegister>(
    ProviderBindings.DATABASE_EVENT_REGISTER,
  );
  await databaseEventRegister.registerDatabaseEventListener();
  
  await webSocketListener.subscribeListener();

  logger.debug('Starting app...');

  app.sequence(MySequence);

  await app.start();

  const url = app.restServer.url;
  logger.info(`Loopback Server is running at ${url}`);
  logger.info(`Socket IO : Watching path : 2`);
  return app;
}
