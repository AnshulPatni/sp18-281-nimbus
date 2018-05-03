import React, {Component} from 'react';
import './mainpage.css';
import Coffee from './Coffee';
import Middle from './Middle';
import Brand from './logo1.png';
import Footer from './Footer';
import axios from 'axios';

import BG from './bg.jpg';

const headers = {
    'Accept': 'application/json'
};




class MainPage extends Component {

  constructor(props) {
      super(props);
      this.state={
        pageView: "coffee",
        dummyData: "lelo mc",
        coffees: [],
        teas: [],
        breads: [],
        desserts: [],
        smoothies: [],
        totalItems: 0
      }

      this.handleCoffeeClick = this.handleCoffeeClick.bind(this);
      this.handleTeaClick = this.handleTeaClick.bind(this);
      this.handleBreadClick = this.handleBreadClick.bind(this);
      this.handleDessertsClick = this.handleDessertsClick.bind(this);
      this.handleSnacksClick = this.handleSnacksClick.bind(this);



    }

    componentWillMount() {

      console.log("Inside ComponentWill Mount of MainPage");

      axios.get('http://13.56.95.132:5000/inventoryCoffees')
     .then((response) => {
       console.log(response.data);
       this.setState({coffees: response.data})

       console.log("checking in mainpage", this.state.coffees);
     })
    .catch((error)=>{
       console.log(error);
    });

    axios.get('http://13.56.95.132:5000/inventoryTeas')
   .then((response) => {
     console.log(response.data);
     this.setState({teas: response.data})

     console.log("checking in mainpage teas", this.state.teas);
   })
  .catch((error)=>{
     console.log(error);
  });


  axios.get('http://54.189.196.198:5000/inventoryBreads')
 .then((response) => {
   console.log(response.data);
   this.setState({breads: response.data})

   console.log("checking in mainpage", this.state.breads);
 })
.catch((error)=>{
   console.log(error);
});

axios.get('http://52.52.176.4:5000/inventoryDesserts')
.then((response) => {
 console.log(response.data);
 this.setState({desserts: response.data})

 console.log("checking in mainpage", this.state.breads);
})
.catch((error)=>{
 console.log(error);
});

}

handleCoffeeClick() {

this.setState({
  pageView: "coffee"
});

}

handleTeaClick() {
  this.setState({
    pageView: "tea"
  });
}

handleBreadClick () {
  this.setState({
    pageView: "breads"
  });
}

handleDessertsClick () {
  this.setState({
    pageView: "desserts"
  });
}

handleSnacksClick () {
  this.setState({
    pageView: "snacks"
  });
}


render() {
  return (
    <div className="mainpage">
        <nav className="navbar navbar-toggleable-md navbar-dark header">
            <a className="navbar-brand" href="#"></a>
            <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="true" aria-label="Toggle navigation" style={{border: "none", outline: "none"}}>
              <span className="navbar-toggler-icon" style={{color: "Tomato", fontSize:"1.5em"}}><i className="fas fa-align-justify"></i></span>
            </button>
            <ul className="navbar-nav mr-auto" >
              <li>
                <span>
                  <img src={Brand} alt="logo" />
                </span>
              </li>
            <div className="collapse navbar-collapse" id="navbarNavAltMarkup">

                    <li className="nav-item active" >
                      <a className="nav-link" href="/">Home</a>
                    </li>
                    <li className="nav-item active" onClick={this.handleCoffeeClick}>
                      <a className="nav-link" href="#">Coffee</a>
                    </li>
                    <li className="nav-item active" onClick={this.handleTeaClick}>
                      <a className="nav-link " href="#">Tea</a>
                    </li>
                    <li className="nav-item active" onClick={this.handleBreadClick}>
                      <a className="nav-link " href="#">Bread</a>
                    </li>
                    <li className="nav-item active" onClick={this.handleDessertsClick}>
                      <a className="nav-link " href="#">Desserts</a>
                    </li>
                    <li className="nav-item active" onClick={this.handleSnacksClick}>
                      <a className="nav-link " href="#">Snacks</a>
                    </li>
                    <li className="nav-item active" onClick={() => {this.props.findStores()}}>
                      <a className="nav-link " href="#">Stores Near You!</a>
                    </li>


            </div>
            </ul>
        </nav>

        <div className="container-fluid content">
            <div className="row">
              <div className="col-lg-1 left">

              </div>
              <div className="col-lg-8 middle">
                  <Middle pageView={this.state.pageView} coffees={this.state.coffees} desserts={this.state.desserts} teas={this.state.teas} breads={this.state.breads} smoothies={this.state.smoothies} addCoffeeToCart={this.props.addCoffeeToCart} addTeaToCart={this.props.addTeaToCart} addBreadToCart={this.props.addBreadToCart} addDessertToCart={this.props.addDessertToCart} addSmoothieToCart={this.props.addSmoothieToCart} incrementLikesCoffee={this.props.incrementLikesCoffee} incrementLikesBread={this.props.incrementLikesBread} incrementLikesDesserts={this.props.incrementLikesDesserts} incrementLikesSmoothies={this.props.incrementLikesSmoothies} incrementLikesTea={this.props.incrementLikesTea} />
              </div>
              <div className="col-lg-3 right">
                Order<br/><br/>
                <h4 className="items-cart-text">Items in the cart: &nbsp;&nbsp;{this.props.itemsCount}</h4><br />
                <button className="btn btn-success" onClick={()=>{this.props.checkCart()}}>Checkout</button>
              </div>
            </div>
        </div>

        <Footer />
    </div>
);
}
  }


export default MainPage;
