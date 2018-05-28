import React, { Component } from 'react';
import logo from './logo.svg';
import StreamWatcher from './StreamWatcher.js';
import AccountStatusForm from './AccountStatusForm.js';
import OrderForm from './OrderForm.js';
import registerServiceWorker from './registerServiceWorker';

class App extends Component {
  render() {
    return (
      <div className="App">
        
        <div>
          <OrderForm />
        </div>

        <div>
          <StreamWatcher />
        </div>

        <div>
        <AccountStatusForm />
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;

