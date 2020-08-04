import { FabricNetworkProvider } from '../providers/fabricNetwork.provider';
import { inject } from '@loopback/core';
import { ProviderBindings } from '../providers/binding.constants';
import { getLogger, Logger } from '../utils/logger';
import { CompagnieAssurance } from '../models';
import { createErrorFromEndorsementError } from '../config/errorMessage';
import { v4 as uuidv4 } from 'uuid';

export class CompagnieAssuranceService {
  private logger: Logger = getLogger(CompagnieAssuranceService.constructor.name);
  constructor(
    @inject(ProviderBindings.FABRIC_NETWORK)
    private fabricNetwork: FabricNetworkProvider,
  ) { }

  async create(compagnieAssurance: CompagnieAssurance){
    this.logger.trace(
      `Create CompagnieAssurance With ${JSON.stringify(compagnieAssurance)}`,
    );
    const transaction = this.fabricNetwork.createTransaction(
      'CreateCompagnieAssurance',
    );
    return transaction
      .submit(
        uuidv4(),
        compagnieAssurance.nom,
        compagnieAssurance.contact,
        compagnieAssurance.adresse,
        // tslint:disable-next-line:no-any
      )
      .catch((error: any) => {
        const elementState = createErrorFromEndorsementError(error);
        return Promise.reject(elementState);
      });
  }

  async read(uuid: string) {
    this.logger.trace('Find CompagnieAssurance With Uuid ' + uuid);

    return this.fabricNetwork
      .evaluateTransaction('GetCompagnieAssuranceByID', uuid);
  }

  /* async update(uuid: string, org: CompagnieAssurance): Promise<Buffer> {
    this.logger.trace(
      `Update CompagnieAssurance With Parameters: 
      uuid: ${uuid},
      nom: ${org.nom},
      contact: ${org.contact},
      adresse: ${org.adresse},`
    );

    const transaction = this.fabricNetwork.createTransaction(
      'UpdateCompagnieAssuranceByID',
    );
    return transaction
      .submit(uuid, org.nom, org.contact, org.adresse )
      .catch((error: any) => {
        const elementState = createErrorFromEndorsementError(error);
        return Promise.reject(elementState);
      });
  }

  async delete(uuid: string) {
    this.logger.trace(`Delete CompagnieAssurance With Uuid  ${uuid}`);

    return this.fabricNetwork
      .createTransaction('UnregisterCompagnieAssuranceByID')
      .submit(uuid)
      .catch((error: any) => {
        const elementState = createErrorFromEndorsementError(error);
        return Promise.reject(elementState);
      });
  }

  async listAllCompagnieAssurances(ObjectType: string) {
    this.logger.trace(`Get all  ${ObjectType}`);

    return this.fabricNetwork
      .evaluateTransaction('GetAllCompagnieAssurances', ObjectType)
      .then(response => Promise.resolve(response));
  }
  
  async getAllCompagnieAssuranceByPage(
    pageSize: string,
    offset: string,
  ) {
    this.logger.trace('Get all by page ');
    return this.fabricNetwork
      .evaluateTransaction(
        'GetAllCompagnieAssurancesByPage',
        pageSize,
        offset,
      )
      .then(response => Promise.resolve(response));
  } */
}
