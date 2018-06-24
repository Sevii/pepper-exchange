import React, { Component } from "react";
import logo from "./logo.svg";
import AccountStatusForm from "./AccountStatusForm.js";
import OrderForm from "./OrderForm.js";
import MarketData from "./MarketData.js";

class App extends Component {
  render() {
    return (
      <div className="App container">
        <div className="car">
          <AccountStatusForm />
        </div>

        <div className="card">
          <OrderForm />
        </div>

        <div className="car">
          <MarketData />
        </div>
      </div>
    );
  }
}

export default App;
