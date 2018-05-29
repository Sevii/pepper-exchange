import React, { Component } from 'react';

class AccountStatusForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      usd: 0,
      btc: 0,
      ltc: 0,
      xmr: 0,
      doge: 0,
      totalValue: 0,
      userId: "BOB"
    };

    this.handleInputChange = this.handleInputChange.bind(this);
    this.getStatus = this.getStatus.bind(this);
  }

  componentDidMount() {
    this.interval = setInterval(() => this.getStatus(), 1000);
  }
  componentWillUnmount() {
    clearInterval(this.interval);
  }


  getStatus(data) {
    var address = 'http://localhost:8080/status/' + this.state.userId
    return fetch(address, {
      method: 'GET'
    })
    .then((responseText) => responseText.json())
    .then((account) =>
      this.setState({
        ["usd"]: account.usd,
        ["btc"]: account.btc,
        ["ltc"]: account.ltc,
        ["doge"]: account.doge,
        ["xmr"]: account.xmr,
        ["totalValue"]: account.totalValue
       })
    )
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.value;
    if(target.name === "selectUser"){
      this.setState({
        ["userId"]: value
       },
        this.getStatus
       );
      
    }
    

  }

  render() {
    return (
      <div className="card">
        <div className="card-body">
          <div className="form-group">
            <form >
              <label>
                <select className="form-control" name="selectUser" value={this.state.userId} onChange={this.handleInputChange}>
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
            </form>
            <div>
              User: {this.state.userId} USD: <bold>{this.state.usd}</bold> Satoshi (BTC/100,000,000): {this.state.btc} LTC: {this.state.ltc} DOGE: {this.state.doge} XMR: {this.state.xmr} Total Value (Satoshi): {this.state.totalValue}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default AccountStatusForm;