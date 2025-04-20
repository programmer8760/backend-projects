import {useState} from 'react';
import './App.css';

function App() {
    const [currentPage, setPage] = useState("length");

    const renderPage = () => {
        switch (currentPage) {
            case "length":
                return (
                    <form>
                        <label className="label">Enter the length to convert</label>
                        <input className="input-field"></input>
                        <label className="label">Unit to convert from</label>
                        <select className="selector">
                            <option value="millimeter">Millimeter</option>
                            <option value="millimeter">Centimeter</option>
                            <option value="millimeter">Meter</option>
                            <option value="millimeter">Kilometer</option>
                            <option value="millimeter">Inch</option>
                            <option value="millimeter">Foot</option>
                            <option value="millimeter">Yard</option>
                            <option value="millimeter">Mile</option>
                        </select>
                        <label className="label">Unit to convert to</label>
                        <select className="selector">
                            <option value="millimeter">Millimeter</option>
                            <option value="millimeter">Centimeter</option>
                            <option value="millimeter">Meter</option>
                            <option value="millimeter">Kilometer</option>
                            <option value="millimeter">Inch</option>
                            <option value="millimeter">Foot</option>
                            <option value="millimeter">Yard</option>
                            <option value="millimeter">Mile</option>
                        </select>
                    </form>
                );
            case "weight":
                return (
                    <form>
                        <label className="label">Enter the weight to convert</label>
                        <input className="input-field"></input>
                    </form>
                );
            case "temperature":
                return (
                    <form>
                        <label className="label">Enter the temperature to convert</label>
                        <input className="input-field"></input>
                    </form>
                );
        }
    };

    return (
        <div id="App">
            <h1 id="title">Unit Converter</h1>
            <div id="topbar" className="topbar">
                <div id="topbar-length" className={currentPage == "length" ? "topbar-button  selected" : "topbar-button"} onClick={() => setPage("length")}>Length</div>
                <div id="topbar-weight" className={currentPage == "weight" ? "topbar-button  selected" : "topbar-button"} onClick={() => setPage("weight")}>Weight</div>
                <div id="topbar-temperature" className={currentPage == "temperature" ? "topbar-button  selected" : "topbar-button"} onClick={() => setPage("temperature")}>Temperature</div>
            </div>
            <div id="workspace" className="workspace">
                {renderPage()}
            </div>
        </div>
    )
}

export default App
