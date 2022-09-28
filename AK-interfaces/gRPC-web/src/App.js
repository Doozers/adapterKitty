import {SsCall, uniCall} from "./proto/proto";
import {useState} from "react";

function App() {
    let utf8Encode = new TextEncoder();

    const [stream, setStream] = useState();
    const [chat, setChat] = useState([]);
    // tmp


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

    if (stream != null) {
        stream.on('data', function (response) {
            setChat(chat => [...chat, String.fromCharCode(...response.getPayload())]);
            //chat.push(String.fromCharCode(...response.getPayload()));
        });

        stream.on('status', function (status) {
            console.log(status);
        });
        stream.on('end', function (end) {
            console.log(end);
        });
    }

    function handleClick2() {
        setStream(SsCall(utf8Encode.encode("3")));
        console.log("I PASS");
    }

    const listItems = chat.map((value) =>
        <li>{value}</li>);

    return (
        <div className="App">
            <header className="App-header">
                <body>
                selem
                </body>
                <button onClick={handleClick2}>
                    LE BOUTTON
                </button>
            </header>
            <ul>{listItems}</ul>
        </div>
    );
}

export default App;
