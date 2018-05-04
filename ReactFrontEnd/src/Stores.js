import React, {Component} from 'react';
import './itempage.css'

import axios from 'axios';


class Stores extends Component {
constructor(props) {
  super(props);
this.state={
  stores: []
}

}

  componentWillMount() {


  }

  render() {
    return (
    <div className="container">
    {this.props.stores.map(store => (
        <div className="stores" style={{marginBottom: 20}}>
          <div className="row coffeerow">
              {store.storeNumber}
          </div>
          <div className="row coffeerow">
              {store.streetNumber}
          </div>
          <div className="row coffeerow">
              {store.street}
          </div>
          <div className="row coffeerow">
              {store.zipcode}
          </div>

        </div>
      ))}
    </div>
  )
  }

}

export default Stores;
