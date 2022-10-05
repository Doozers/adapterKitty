import proto from "./adapterkit_pb";

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

function SsCall(value) {
    let request = new proto.AdapterRequest();
    request.setPayload(value);
    return client.serverStreamingAdapter(request, {});
}

export {
    uniCall,
    SsCall
};