import { inject } from "@loopback/core";
import { ProviderBindings, FabricNetworkProvider } from ".";

export class WebSocketListener {
    constructor(
        @inject(ProviderBindings.FABRIC_NETWORK)
        private fabricNetwork: FabricNetworkProvider,
        @inject(ProviderBindings.SOCKET_IO_SERVER)
        private socketServer: SocketIO.Server,
    ) { }

    async subscribeListener() {
        const listener = await this.fabricNetwork.addEventListener('.*', async (
            error: Error,
            event: any,
            block_num: any,
            txnid: any,
            status: any,
        ) => {
                const obj = {
                    name: event.event_name,
                    payload: event.payload.toString(),
                    txtId: txnid,
                }
                const outputString: string = JSON.stringify(obj);

                this.socketServer.emit('log', outputString);
        });
    }
}