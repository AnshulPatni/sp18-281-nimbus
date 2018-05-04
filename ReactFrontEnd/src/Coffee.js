import React, {Component} from 'react';
import './itempage.css';

class Coffee extends Component {

constructor(props) {
      super(props);
this.state={
  coffees: [
    {
      cost: 5.2,
      item: "latte",
      likes: 0,
      status: 0,
      typeService: "coffee"
    },
    {
      cost: 4.0,
      item: "mocha",
      likes: 0,
      status: 0,
      typeService: "coffee"
    }
  ]
}
}

componentWillMount() {

console.log("lets check",this.props.coffees);
this.setState({
  coffees: this.props.coffees
})

}

addCoffeeToCart(data) {

  console.log("Inside AddCoffeeToCart");

  fetch('http://52.53.240.238:5000/addToCartCoffees?item='+data.item, {
      method: 'GET',
      headers: {
          'Content-Type': 'application/json'
      },
      credentials: "include"
  }).then(res => {
      console.log(res);
      return res.status;
  })
      .catch(error => {
          console.log("This is error");
          return error;
      });

}

render() {
  return (
    <div className="coffeepage">
        <div className="container-fluid">
          <div className="row coffeerow" style={{backgroundColor: "#282726"}}>
              <div className="col-lg-2 desc">
                Likes
              </div>
              <div className="col-lg-4 desc">
                Name
              </div>
              <div className="col-lg-3 desc">
                Price
              </div>
              <div className="col-lg-3 desc">

              </div>
          </div>

{this.props.coffees.map(coffee => (
          <div className="row coffeerow" style={{marginTop: 3}}>
              <div className="col-lg-2 desc">
                {coffee.likes}&nbsp; <i className="fas fa-heart fa-1x" style={{color: "red", fontSize: 18}} onClick={() => {this.props.incrementLikesCoffee({item: coffee.item, likes: coffee.likes})}}></i>
              </div>
              <div className="col-lg-4 desc">
                {coffee.item}
              </div>
              <div className="col-lg-3 desc">
                {coffee.cost}
              </div>
              <div className="col-lg-3 desc" style={{paddingTop: 5}}>
                <button type="button" className="btn btn-primary" onClick={()=>{this.props.addCoffeeToCart({item: coffee.item})}}>Add to Cart </button>
              </div>
          </div>
))}





        </div>
        <br/>

    </div>
);
}
  }


export default Coffee;
