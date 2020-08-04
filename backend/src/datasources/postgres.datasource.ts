import {inject} from '@loopback/core';
import {juggler} from '@loopback/repository';

export class PostgresDataSource extends juggler.DataSource {
  static dataSourceName = 'postgres';

  constructor(
    @inject('datasources.config.sqlstore', {optional: true})
    dataSourceConfig: object,
  ) {
    super(dataSourceConfig);
  }
}
