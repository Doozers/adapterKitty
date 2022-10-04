import proto from "./adapterkit_pb";

const {AdapterKitServicePromiseClient} = require('./adapterkit_grpc_web_pb.js');

let client = new AdapterKitServicePromiseClient('http://127.0.0.1:9315');

function uniCall(value) {
    let request = new proto.AdapterRequest();
    request.setPayload(value);
    return client.uniDirectionalAdapter(request, {});
}

function SsCall(value) {
    let request = new proto.AdapterRequest();
    request.setPayload(value);
    return client.serverStreamingAdapter(request, {});
}

export {
    uniCall,
    SsCall
};