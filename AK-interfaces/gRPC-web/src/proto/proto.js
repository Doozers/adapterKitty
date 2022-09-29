import proto from "./adapterkit_pb";
import custom from "./announcement_pb";

const {AdapterKitServiceClient} = require('./adapterkit_grpc_web_pb.js');

let client = new AdapterKitServiceClient('http://127.0.0.1:9315');

function uniCall(value) {
    let request = new proto.AdapterRequest();
    request.setPayload(value);
    return client.uniDirectionalAdapter(request, {}, function (err, response) {
        if (err) {
            console.log(err);
        } else {
            return response;
        }
    });
}

function SsCall() {
    let sheme = new custom.ConnectionRequest();
    sheme.setAsktoconnect(true);

    let request = new proto.AdapterRequest();
    request.setPayload(sheme.serializeBinary());
    return client.serverStreamingAdapter(request, {});
}

export {
    uniCall,
    SsCall
};