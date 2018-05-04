import React, {Component} from 'react';
import './Cart.css';

import axios from 'axios';


class Cart extends Component {
constructor(props) {
  super(props);
this.state={
  items: []
}

}



  componentWillMount() {
    console.log("Inside ComponentWill mount of cart");

    axios.get('http://13.56.95.132:5000/cartItemsCoffees')
   .then((response) => {
     console.log(response.data);
     var coffees = response.data;
     var oldItems = this.state.items;
     var newItems = oldItems.concat(coffees);
     this.setState({items: newItems})

     console.log("checking in cart items", this.state.coffees);
   })

  .catch((error)=>{
     console.log(error);
  });

  axios.get('http://54.189.196.198:5000/cartItemsBreads')
 .then((response) => {
   console.log(response.data);
   var breads = response.data;
   var oldItems = this.state.items;
   var newItems = oldItems.concat(breads);
   this.setState({items: newItems})

   console.log("checking in cart items", this.state.items);
 })
.catch((error)=>{
   console.log(error);
});

axios.get('http://52.52.176.4:5000/cartItemsDesserts')
.then((response) => {
 console.log(response.data);
 var desserts = response.data;
 var oldItems = this.state.items;
 var newItems = oldItems.concat(desserts);
 this.setState({items: newItems})

 console.log("checking in cart items", this.state.items);
})
.catch((error)=>{
 console.log(error);
});


  }

  render() {
    return (
    <div className="container">
        <div className="row">
          <h1>  Items in the cart:</h1>
        </div>
        <br />
        <div className="row coffeerow" style={{marginBottom: 10, backgroundColor: "gray", textAlign: "center"}}>
            <div className="col-lg-6">
              Item Name
            </div>
            <div className="col-lg-6">
              Item Cost
            </div>
        </div>
        {this.state.items.map(item => (

          <div className="row coffeerow">
              <div className="col-lg-6">
                {item.item}
              </div>
              <div className="col-lg-6">
                {item.cost}
              </div>
          </div>
        ))}

        <br/>
        <br/>
        <div className="row coffeerow">
            <button className="btn btn-success" onClick={() => {this.props.processOrder()}}>Process Order</button>
        </div>
    </div>
  )
  }

}


export default Cart;
