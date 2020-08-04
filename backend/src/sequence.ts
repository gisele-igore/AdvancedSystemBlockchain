import {inject} from '@loopback/context';
import {
  FindRoute,
  InvokeMethod,
  ParseParams,
  Reject,
  RequestContext,
  RestBindings,
  Send,
  SequenceHandler,
} from '@loopback/rest';
import * as uuid from 'uuid';
import {getLogger, Logger} from './utils/logger';

const SequenceActions = RestBindings.SequenceActions;

export class MySequence implements SequenceHandler {
  private logger: Logger = getLogger(MySequence.constructor.name);
  constructor(
    @inject(SequenceActions.FIND_ROUTE) protected findRoute: FindRoute,
    @inject(SequenceActions.PARSE_PARAMS) protected parseParams: ParseParams,
    @inject(SequenceActions.INVOKE_METHOD) protected invoke: InvokeMethod,
    @inject(SequenceActions.SEND) public send: Send,
    @inject(SequenceActions.REJECT) public reject: Reject,
  ) {}

  async handle(context: RequestContext) {
    const requestId: string = uuid.v4();
    try {
      const {request, response} = context;
      context.bind<string>('requestId').to(requestId);
      request.url = request.url.replace('api/', '');
      this.logger.debug(
        `[${requestId}] Request received with method[${request.method}] url[${request.url}] body[${request.body}] params[${request.params}]`,
      );
      const route = this.findRoute(request);
      const args = await this.parseParams(request, route);
      const result = await this.invoke(route, args);
      this.send(response, result);
      this.logger.debug(
        `[${requestId}] Response sent with statusCode[${response.statusCode}] statusMessage[${response.statusMessage}]`,
      );
    } catch (err) {
      this.logger.error(`[${requestId}] Error handling request : ${err}`);
      this.reject(context, err);
    }
  }
}
