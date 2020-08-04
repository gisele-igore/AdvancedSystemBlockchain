import {Entity, property, model} from '@loopback/repository';
@model()
export class LogsData extends Entity {
  @property({
    type: 'string',
    id: true,
    generated: true,
  })
  uuid: string;
  @property({
    type: 'string',
  })
  contents: string;

  @property({
    type: 'string',
  })
  timestamp: Number;
  constructor(data?: Partial<LogsData>) {
    super(data);
  }
}
