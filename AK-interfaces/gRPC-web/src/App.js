import {SsCall, uniCall} from "./proto/proto";
import {useState} from "react";

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

        setStream(SsCall(utf8Encode.encode()));
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
                selem
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
