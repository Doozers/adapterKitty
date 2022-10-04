import {useReducer, useState} from "react";
import custom from "./proto/occurrence_pb"
import {uniCall} from "./proto/proto";

function App() {

    const formReducer = (state, event) => {
        return {
            ...state,
            [event.name]: event.value
        }
    }

    const [formData, setFormData] = useReducer(formReducer, {});

    const handleChange = event => {
        setFormData({
            name: event.target.name,
            value: event.target.value,
        });
    }

    function handleOnClick() {
        let sheme = new custom.wikiRequest();
        sheme.setKeyword(formData.keyword)
        sheme.setNeedle(formData.needle)

        uniCall(sheme.serializeBinary()).then((response, err) => {

            if (err) {
                console.log(err);
            } else {
                let shemeRes = custom.wikiResponse.deserializeBinary(response.getPayload())
                console.log(shemeRes.getOccurrence())

                alert("there is "+shemeRes.getOccurrence()+" occurrences of "+formData.needle+" in the webpage of "+formData.keyword)
            }
        })
    }

    return (
        <div className="App">
            <header className="App-header">
                <body>
                selem
                </body>
                <form >
                    <fieldset>
                        <label>
                            <p>Keyword</p>
                            <input name="keyword" onChange={handleChange} step="1"/>
                        </label>
                        <label>
                            <p>Needle</p>
                            <input name="needle" onChange={handleChange}/>
                        </label>
                    </fieldset>
                </form>
                <button onClick={handleOnClick}>
                    LE BOUTTON
                </button>
            </header>
        </div>
    );
}

export default App;
