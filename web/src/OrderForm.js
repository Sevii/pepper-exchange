import React, { Component } from 'react';

class OrderForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      price: 100,
      number: 10,
      exchange: "BTCUSD",
      direction: "ask",
      userId: "Bob",
      cancelId: "None"
    };

    this.handleInputChange = this.handleInputChange.bind(this);
    this.submitOrderRequest = this.submitOrderRequest.bind(this);
  }

  submitOrderRequest(event) {
    fetch('http://localhost:8080/order', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(this.state)
    })
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.value;
    if(target.name === "numberOfCoins"){
      this.setState({
      ["number"]: value
    });
    }
    if(target.name === "pricePerCoin"){
      this.setState({
      ["price"]: value
    }); 
    }
    if(target.name === "selectExchange"){
      this.setState({
      ["exchange"]: value
    });
    }
    if(target.name === "selectDirection"){
      this.setState({
      ["direction"]: value
    });
    }
    if(target.name === "selectUser"){
      this.setState({
      ["userId"]: value
    }); 
    }
    if(target.name === "cancelId"){
      this.setState({
      ["cancelId"]: value
    }); 
    }


  }

  render() {
    return (
      <form>
        <label>
          Number of coins:
          <input
            name="numberOfCoins"
            type="number"
            value={this.state.number}
            onChange={this.handleInputChange} />
        </label>
        <label>
          Price per coin:
          <input
            name="pricePerCoin"
            type="number"
            value={this.state.price}
            onChange={this.handleInputChange} />
        </label>
      <label>
          Pick your exchange.
          <select name="selectExchange" value={this.state.exchange} onChange={this.handleInputChange}>
            <option value="BTCUSD">Trade bitcoin for USD</option>
            <option value="BTCLTC">Trade bitcoin for LTC</option>
            <option value="BTCDOGE">Trade bitcoin for DOGE</option>
            <option value="BTCXMR">Trade bitcoin for XMR</option>
          </select>
        </label>
       <label>
          Pick your direction (Bid/Ask/Cancel).
          <select name="selectDirection" value={this.state.direction} onChange={this.handleInputChange}>
            <option value="bid">Bid for a coin</option>
            <option value="ask">Ask for a price</option>
            <option value="cancel">Cancel an order</option>
          </select>
        </label>
        <label>
          Your Name
          <select name="selectUser" value={this.state.userId} onChange={this.handleInputChange}>
            <option value="BOB">BOB</option>
            <option value="ALICE">ALICE</option>
            <option value="ROBODOG">ROBODOG</option>
            <option value="KID1">KID1</option>
            <option value="KID2">KID2</option>
            <option value="KID3">KID3</option>
            <option value="KID4">KID4</option>
            <option value="OTHERKID">OTHERKID</option>
          </select>
        </label>
         <label>
          Order to cancel ID
          <input
            name="cancelId"
            type="text"
            value={this.state.cancelId}
            onChange={this.handleInputChange} />
        </label>
        <label>
          Submit
          <button
            name="submitOrder"
            type="button"
            value="Submit"
            onClick={this.submitOrderRequest} />
        </label>
      </form>
    );
  }
}

export default OrderForm;