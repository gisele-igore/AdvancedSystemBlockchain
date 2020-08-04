import * as log4js from 'log4js';
import {configurationGeneral} from './../config';

export {Logger} from 'log4js';

export function getLogger(category?: string | undefined): log4js.Logger {
  return log4js.getLogger(category);
}

export function configureLogger() {
  log4js.configure({
    appenders: {
      console: {
        type: 'console',
        layout: {
          type: 'pattern',
          pattern: `%[${configurationGeneral.logger.pattern}%]`,
        },
      },
    },
    categories: {
      default: {
        appenders: ['console'],
        level: configurationGeneral.logger.level,
      },
    },
  });
}

export function shutdownLogger(): void {
  log4js.shutdown(() => {});
}
