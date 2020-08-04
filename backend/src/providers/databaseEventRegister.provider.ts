import {inject} from '@loopback/core';
import {getLogger, Logger} from '../utils/logger';
import { FabricNetworkProvider } from './fabricNetwork.provider';
import { ProviderBindings } from './binding.constants';
import {LogsData} from '../models';
import {LogsDataRepository} from '../repositories';
import {repository} from '@loopback/repository';

export class DatabaseEventRegister {
    private logger: Logger = getLogger(FabricNetworkProvider.constructor.name);  
    constructor(
    @inject(ProviderBindings.FABRIC_NETWORK)
    private fabricNetwork: FabricNetworkProvider,
    @repository(LogsDataRepository)
    private logsDataRepository: LogsDataRepository,
    ) { }

    async registerDatabaseEventListener() {
        const listenerUuid = await this.fabricNetwork.addEventListener(
            '.*',
            async (
              error: Error,
              event: any,
              block_num: any,
              txnid: any,
              status: any,
            ) => {
               const obj = {
                        name: event.event_name,
                        payload: event.payload.toString(),
                        txId: txnid,
                      };
        
                const outputString: string = JSON.stringify(obj);
                await this.storeNewLogInDatabase(outputString);

                this.logger.debug('\n' + outputString + '\n');                
            },
        );
    }

    async storeNewLogInDatabase(contents: string): Promise<LogsData> {
        const date = new Date();
        const timestamp = date.getTime();    
        return await this.logsDataRepository.create({
          contents,
          timestamp,
        });
      }

}