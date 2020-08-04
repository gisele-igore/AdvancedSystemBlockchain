import {DefaultCrudRepository} from '@loopback/repository';
import {LogsData} from '../models';
import {PostgresDataSource} from '../datasources';
import {inject} from '@loopback/core';

export class LogsDataRepository extends DefaultCrudRepository<
  LogsData,
  typeof LogsData.prototype.uuid
> {
  constructor(@inject('datasources.postgres') dataSource: PostgresDataSource) {
    super(LogsData, dataSource);
  }
}
