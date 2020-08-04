import { post, requestBody, get, del, param, put, } from '@loopback/rest';
import { CompagnieAssuranceService, ServiceBindings } from '../services';
import { HttpErrors } from '@loopback/rest';
import { CompagnieAssurance } from '../models';
import { inject } from '@loopback/core';
import { getLogger, Logger } from '../utils/logger';
import { response204, response409, response200, response404, response200_1, jsonFormatter } from '../config/restDecoration';
import { ErrorMessage } from '../config/errorMessage';

const examples = {
  "nom": "SCHAIN",
  "contact": "Scorechain",
  "adresse": "ISP"
}
const getExample: CompagnieAssurance = JSON.parse(`{
  "uuid": "263126AA-1BAF-4F34-85EC-739399BC1054",
  "nom": "SCHAIN",
  "contact": "Scorechain",
  "adresse": "ISP"
}`)

export class CompagnieAssuranceController {
  private logger: Logger = getLogger(CompagnieAssuranceController.constructor.name);
  constructor(
    @inject(ServiceBindings.COMPAGNIE_ASSURANCE_SERVICE)
    private compagnieAssuranceService: CompagnieAssuranceService,
  ) { }

  @post('/compagnieAssurance', {
    responses: {
      '200': response200(CompagnieAssurance, 'CompagnieAssurance', getExample),
      '409': response409(CompagnieAssurance)
    },
  })
  async create(
    @requestBody(
      {
        description: 'CompagnieAssurance object that needs to be added to the network',
        required: true,
        content: {
          'application/json': {
            'Content-Type': { 'x-ts-type': CompagnieAssurance },
            example: examples,
          },
        },
      }) compagnieAssurance: CompagnieAssurance) {
    return this.compagnieAssuranceService
      .create(compagnieAssurance)
      .then(compagnieAssurance => JSON.parse(compagnieAssurance.toString()))
      .catch(err => {
        if (err.message === ErrorMessage.COMPAGNIE_ASSURANCE_ALREADY_EXIST){
          throw new HttpErrors.Conflict(err.message);
        }
      });
  }

  @get('/compagnieAssurance/{compagnieAssuranceUuid}', {
    responses: {
      '200': response200_1(CompagnieAssurance, 'CompagnieAssurance', getExample),
      '404': response404('CompagnieAssurance'),
    },
  })
  async getOrganization(@param.path.string('compagnieAssuranceUuid') idOrganization: string) {
    return this.compagnieAssuranceService
      .read(idOrganization)
      .then(compagnieAssurance => JSON.parse(compagnieAssurance.toString()))
      .catch(err => {
        throw new HttpErrors.NotFound(err.message);
      });
  }

}
 