import proto from "./proto/adapterkit_pb";

const {AdapterKitServiceClient} = require('./proto/adapterkit_grpc_web_pb.js');

function App() {
    let client = new AdapterKitServiceClient('http://127.0.0.1:9315');

    function handleOnClick() {
        let request = new proto.AdapterRequest();
        client.uniDirectionalAdapter(request, {}, function(err, response) {
            if (err) {
                console.log(err);
            }  else {
                console.log(String.fromCharCode(...response.getPayload()));
            }
        });
    }

    return (
        <div className="App">
            <header className="App-header">
                <body>
                selem
                </body>
                <button onClick={handleOnClick}>
                    LE BOUTTON
                </button>
            </header>
        </div>
    );
}

export default App;
