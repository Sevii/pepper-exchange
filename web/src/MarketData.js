import React, { Component } from "react";

class MarketData extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      ltcPrice: 0,
      usdPrice: 0,
      xmrPrice: 0,
      dogePrice: 0,
      fills: []
    };

    this.getMarketData = this.getMarketData.bind(this);
  }
  componentDidMount() {
    this.interval = setInterval(() => this.getMarketData(), 1000);
  }
  componentWillUnmount() {
    clearInterval(this.interval);
  }

  getMarketData(data) {
    var address = "http://localhost:8080/marketdata";
    return fetch(address, {
      method: "GET"
    })
      .then(responseText => responseText.json())
      .then(account => {
        this.setState({
          ["usdPrice"]: account.usdPrice,
          ["ltcPrice"]: account.ltcPrice,
          ["dogePrice"]: account.dogePrice,
          ["xmrPrice"]: account.xmrPrice,
          ["fills"]: account.lastFills
        });
      });
  }

  render() {
    const listItems = this.state.fills.map(fill => (
      <li className="list-group-item" key={fill.id}>
        Fill: {fill.id} Exchange: {fill.exchange} Price: {fill.price} Number:{" "}
        {fill.number} Timestamp: {fill.timestamp}
      </li>
    ));

    return (
      <div>
        <div className="card">
          <div className="card-body">
            Latest Prices: Bitcoin per Dollar ($): {this.state.usdPrice} Bitcoin
            per Litecoin (LTC): {this.state.ltcPrice} Bitcoin per Dogecoin
            (DOGE): {this.state.dogePrice} Bitcoin per Monero (XMR):{" "}
            {this.state.xmrPrice}
          </div>
        </div>

        <div>
          <ul className="list-group">{listItems}</ul>
        </div>
      </div>
    );
  }
}

export default MarketData;
