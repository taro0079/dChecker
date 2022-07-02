import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {ShowResult} from '../wailsjs/go/main/App';

function App() {
    const [resultText, setResultText] = useState("免許証番号を入れて👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    function checkdigit() {
        ShowResult(name).then(updateResultText);
    }

    return (
        <div id="App">
            <div className='text-3xl'>DigitChecker</div>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={checkdigit}>Check</button>
            </div>
        </div>
    )
}

export default App
