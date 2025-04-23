import { useState } from 'react';
import './App.css';
import { ConvertLength, ConvertWeight, ConvertTemperature } from "../wailsjs/go/main/App";

function App() {
    const [currentPage, setPage] = useState("length");
    const [result, setResult] = useState(null);

    let convertedValue;

    const convertLength = (e) => {
        e.preventDefault();

        const formData = new FormData(e.target);
        const length = parseFloat(formData.get("length"));
        const from = parseInt(formData.get("from"));
        const to = parseInt(formData.get("to"));
        ConvertLength(length, from, to).then((response) => {
            setResult(response);
        });
    }

    const convertWeight = (e) => {
        e.preventDefault();

        const formData = new FormData(e.target);
        const weight = parseFloat(formData.get("weight"));
        const from = parseInt(formData.get("from"));
        const to = parseInt(formData.get("to"));
        ConvertWeight(weight, from, to).then((response) => {
            setResult(response);
        });
    }

    const convertTemperature = (e) => {
        e.preventDefault();
        
        const formData = new FormData(e.target);
        const temperature = parseFloat(formData.get("temperature"));
        const from = parseInt(formData.get("from"));
        const to = parseInt(formData.get("to"));
        ConvertTemperature(temperature, from, to).then((response) => {
            setResult(response);
        });
    }

    const reset = () => {
        setResult(null);
    }

    const renderForm = () => {
        if(result != null) {
            return;
        }
        switch (currentPage) {
            case "length":
                return (
                    <form onSubmit={convertLength} id="length-form">
                        <label className="label">Enter the length to convert</label>
                        <input name="length" className="input-field"></input>
                        <label className="label">Unit to convert from</label>
                        <select name="from" className="selector" form="length-form">
                            <option value="0">Millimeter</option>
                            <option value="1">Centimeter</option>
                            <option selected="selected" value="2">Meter</option>
                            <option value="3">Kilometer</option>
                            <option value="4">Inch</option>
                            <option value="5">Foot</option>
                            <option value="6">Yard</option>
                            <option value="7">Mile</option>
                        </select>
                        <label className="label">Unit to convert to</label>
                        <select name="to" className="selector" form="length-form">
                            <option value="0">Millimeter</option>
                            <option value="1">Centimeter</option>
                            <option selected="selected" value="2">Meter</option>
                            <option value="3">Kilometer</option>
                            <option value="4">Inch</option>
                            <option value="5">Foot</option>
                            <option selected="selected" value="6">Yard</option>
                            <option value="7">Mile</option>
                        </select>
                        <button type="submit" className="button">Convert</button>
                    </form>
                );
            case "weight":
                return (
                    <form onSubmit={convertWeight}  id="weight-form">
                        <label className="label">Enter the weight to convert</label>
                        <input name="weight" className="input-field"></input>
                        <label className="label">Unit to convert from</label>
                        <select name="from" className="selector" form="weight-form">
                            <option value="0">Milligram</option>
                            <option value="1">Gram</option>
                            <option selected="selected" value="2">Kilogram</option>
                            <option value="3">Ounce</option>
                            <option value="4">Pound</option>
                        </select>
                        <label className="label">Unit to convert to</label>
                        <select name="to" className="selector" form="weight-form">
                            <option value="0">Milligram</option>
                            <option value="1">Gram</option>
                            <option value="2">Kilogram</option>
                            <option value="3">Ounce</option>
                            <option selected="selected" value="4">Pound</option>
                        </select>
                        <button type="submit" className="button">Convert</button>
                    </form>
                    
                );
            case "temperature":
                return (
                    <form onSubmit={convertTemperature} id="temperature-form">
                        <label className="label">Enter the temperature to convert</label>
                        <input name="temperature" className="input-field"></input>
                        <label className="label">Unit to convert from</label>
                        <select name="from" className="selector" form="temperature-form">
                            <option selected="selected" value="0">Celsius</option>
                            <option value="1">Fahrenheit</option>
                            <option value="2">Kelvin</option>
                        </select>
                        <label className="label">Unit to convert to</label>
                        <select name="to" className="selector" form="temperature-form">
                            <option value="0">Celsius</option>
                            <option selected="selected" value="1">Fahrenheit</option>
                            <option value="2">Kelvin</option>
                        </select>
                        <button type="submit" className="button">Convert</button>
                    </form>
                );
            default: 
                return (
                    <p>?</p>
                )
        }
    };

    const renderResult = () => {
        if (result != null) {
            return (
                <div>
                    <div className="label">Result of your calculation</div>
                    <div className="big-text">{result}</div>
                    <button onClick={reset} className="button">Reset</button>
                </div>
            );

        }
    }

    return (
        <div id="App">
            <h1 className="big-text">Unit Converter</h1>
            <div id="topbar" className="topbar">
                <div id="topbar-length" className={currentPage == "length" ? "topbar-button  selected" : "topbar-button"} onClick={() => {setPage("length"); setResult(null)}}>Length</div>
                <div id="topbar-weight" className={currentPage == "weight" ? "topbar-button  selected" : "topbar-button"} onClick={() => {setPage("weight"); setResult(null)}}>Weight</div>
                <div id="topbar-temperature" className={currentPage == "temperature" ? "topbar-button  selected" : "topbar-button"} onClick={() => {setPage("temperature"); setResult(null)}}>Temperature</div>
            </div>
            <div id="workspace" className="workspace">
                {renderForm()}
                {renderResult()}
            </div>
        </div>
    )
}

export default App
