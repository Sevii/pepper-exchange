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
      userId: "BOB"
    };

    this.handleInputChange = this.handleInputChange.bind(this);
    this.getStatus = this.getStatus.bind(this);
  }


 componentDidMount() {
    this.getStatus()
  }


  getStatus(data) {
    var address = 'http://localhost:8080/status/' + this.state.userId
    return fetch(address, {
      method: 'GET'
    })
    .then((responseText) => responseText.json())
    .then((account) =>
      this.setState({
        ["usd"]: account.USD,
        ["btc"]: account.BTC,
        ["ltc"]: account.LTC,
        ["doge"]: account.DOGE,
        ["xmr"]: account.XMR,
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
      <div>
        <form>
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
        </form>
        <div>
          User: {this.state.userId} USD: {this.state.usd} BTC: {this.state.btc} LTC: {this.state.ltc} DOGE: {this.state.doge} XMR: {this.state.xmr}
        </div>
      </div>
    );
  }
}

export default AccountStatusForm;