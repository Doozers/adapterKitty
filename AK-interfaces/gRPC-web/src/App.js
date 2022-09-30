import {SsCall, uniCall} from "./proto/proto";
import {useState} from "react";

import custom from "./proto/announcement_pb";


function App() {
    let utf8Encode = new TextEncoder();

    const [stream, setStream] = useState(null);
    const [chat, setChat] = useState([]);


    function handleClick1() {
        uniCall(utf8Encode.encode("hello")).then((response, err) => {
            if (err) {
                console.log("I PASS");
                console.log(err);
            } else {
                console.log(String.fromCharCode(...response.getPayload()));
            }
        })

    }


    function handleClick2() {

        let sheme = new custom.ConnectionRequest();
        sheme.setAsktoconnect(true);

        setStream(SsCall(sheme.serializeBinary()));
        if (stream != null) {
            stream.on('data', function (response) {
                setChat(chat => [...chat, String.fromCharCode(...response.getPayload())]);
            });
            stream.on('status', function (status) {
                console.log("status: ", status);
            });
            stream.on('end', function (end) {
                console.log("end: ", end);
            });
        }
    }

    const listItems = chat.map((value) =>
        <li>{value}</li>);

    return (
        <div className="App">
            <header className="App-header">
                <body>
                    Selem
                </body>
                <button onClick={handleClick2}>
                    Connect
                </button>
            </header>
            <ul>{listItems}</ul>
        </div>
    );
}

export default App;
