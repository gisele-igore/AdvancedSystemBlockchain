import {Entity, model, property} from '@loopback/repository';

@model()
export class CompagnieAssurance extends Entity {
  @property({
    type: 'string',
    id: true,
  })
  uuid: string;

  @property({
    type: 'string',
  })
  nom: string;

  @property({
    type: 'string',
  })
  contact: string;

  @property({
    type: 'string',
  })
  adresse: string;

  constructor(data?: Partial<CompagnieAssurance>) {
    super(data);
  }
}
