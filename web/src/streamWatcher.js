import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Websocket from 'react-websocket';
import _ from 'lodash';
import 'whatwg-fetch'

class StreamWatcher extends Component {
    constructor(props) {
      super(props);
      this.state = {
        fills: []
      };
    }

 
    handleData(data) {
   
      let lines = data.split(/\n/)

      // console.log(lines[0])
      let latest = _.takeRight(lines, 10).filter(Boolean)
      
      // let objs = JSON.parse(latest[1])
      let objs = _.map(latest, line => JSON.parse(line.trim()))

      // console.log(objs)
      this.setState({fills: objs });

    }
 


    render() {

      const listItems = this.state.fills.map((fill) =>
        <li key={fill.id}>
          Fill: {fill.id} Exchange: {fill.exchange} Direction {fill.direction} Price: {fill.price} Number: {fill.number} Timestamp: {fill.timestamp}
        </li>
      );


      return (
        <div>
 
          <Websocket url='ws://localhost:8080/stream/fills/BTCUSD'
              onMessage={this.handleData.bind(this)}/>
              <ul>{listItems}</ul>

        </div>
      );
    }
  }

export default StreamWatcher;

    function Item(props) {
      return <li>{props}</li>;
    }

    function TodoList(props) {
      return (
        <ul>
          {props.map((fill) => <Item key={fill.id} message={fill.exchange} />)}
        </ul>
      );
    }
