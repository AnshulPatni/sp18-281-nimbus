import React, {Component} from 'react';
import {Route, withRouter} from 'react-router-dom';
import Cart from './Cart';
import MainPage from './MainPage';
import Stores from './Stores';
import Home from './Home';
import axios from 'axios';


class Routing extends Component {
  constructor(props) {
    super(props);

  this.state={
    totalItems: 0,
    stores:[]
  }

}
checkCart = () => {
      this.props.history.push('/cart');
  };

incrementItemsCart = () => {
        var totalItemsNew = this.state.totalItems;
        totalItemsNew = totalItemsNew + 1;

        this.setState({
          totalItems: totalItemsNew
        });
};

processOrder = () => {
  console.log("Inside processOrder");

  axios.get('http://13.56.95.132:5000/processOrdersCoffees')
 .then((response) => {
   console.log(response.data);

   // send to Confirmation page.
   // send email to a sample user

   this.props.history.push('/');
   var tot = 0;
   this.setState({totalItems: tot})

 })
.catch((error)=>{
   console.log(error);
});

axios.get('http://52.52.176.4:5000/processOrdersDesserts')
.then((response) => {
 console.log(response.data);

 // send to Confirmation page.
 // send email to a sample user

 this.props.history.push('/');
 var tot = 0;
 this.setState({totalItems: tot})

})
.catch((error)=>{
 console.log(error);
});

axios.get('http://54.189.196.198:5000/processOrdersBreads')
.then((response) => {
 console.log(response.data);

 // send to Confirmation page.
 // send email to a sample user

 this.props.history.push('/');
 var tot = 0;
 this.setState({totalItems: tot})

})
.catch((error)=>{
 console.log(error);
});

}

findStores = () => {

  console.log("Inside FindStores");

  axios.get('http://13.56.95.132:5000/findStores?zipcode=95126')
  .then((response) => {
   console.log(response.data);
   this.setState({stores: response.data})
   this.props.history.push('/stores')
  })
  .catch((error)=>{
   console.log(error);
  });

}

incrementLikesCoffee = (data) => {

  console.log("Inside incrementLikes");

  axios.get('http://13.56.95.132:5000/likeIncreaseCoffees?item='+data.item+'&likes='+data.likes)
  .then((response) => {
   console.log(response.data);
   var dummyStores = [];
   this.setState({
     stores: dummyStores
   });
   this.props.history.push('/');
  })
  .catch((error)=> {
   console.log(error);
  });

}

incrementLikesTea = (data) => {

  console.log("Inside incrementLikes Tea");

  axios.get('http://13.56.95.132:5000/likeIncreaseCoffees?item='+data.item+'&likes='+data.likes)
  .then((response) => {
   console.log(response.data);
   var dummyStores = [];
   this.setState({
     stores: dummyStores
   });
   this.props.history.push('/');
  })
  .catch((error)=> {
   console.log(error);
  });

}

incrementLikesBread = (data) => {

  console.log("Inside incrementLikes");

  axios.get('http://54.189.196.198:5000/likeIncreaseBreads?item='+data.item+'&likes='+data.likes)
  .then((response) => {
   console.log(response.data);
   var dummyStores = [];
   this.setState({
     stores: dummyStores
   });
   this.props.history.push('/');
  })
  .catch((error)=> {
   console.log(error);
  });
}

incrementLikesDesserts = (data) => {

  console.log("Inside incrementLikes of Desserts");

  axios.get('http://52.52.176.4:5000/likeIncreaseDesserts?item='+data.item+'&likes='+data.likes)
  .then((response) => {
   console.log(response.data);
   var dummyStores = [];
   this.setState({
     stores: dummyStores
   });
   this.props.history.push('/');
  })
  .catch((error)=> {
   console.log(error);
  });
}

incrementLikesSmoothies = (data) => {

  console.log("Inside incrementLikes of Smoothies");

  axios.get('http://52.52.176.4:5000/likeIncreaseDesserts?item='+data.item+'&likes='+data.likes)
  .then((response) => {
   console.log(response.data);
   var dummyStores = [];
   this.setState({
     stores: dummyStores
   });
   this.props.history.push('/');
  })
  .catch((error)=> {
   console.log(error);
  });
}


addCoffeeToCart = (data) => {
        // Backend Api to feed into the database

        console.log("Inside AddCoffeeToCart", data);

        axios.get('http://13.56.95.132:5000/addToCartCoffees?item='+data.item)
       .then((response) => {
         console.log(response.data);

       })
      .catch((error)=>{
         console.log(error);
      });

        this.incrementItemsCart();
};

addBreadToCart = (data) => {
        // Backend Api to feed into the database

        console.log("Inside AddBreadToCart", data);

        axios.get('http://54.189.196.198:5000/addToCartBreads?item='+data.item)
       .then((response) => {
         console.log(response.data);

       })
      .catch((error)=>{
         console.log(error);
      });

        this.incrementItemsCart();
};

addDessertToCart = (data) => {
        // Backend Api to feed into the database

        console.log("Inside addDessertToCart", data);

        axios.get('http://52.52.176.4:5000/addToCartDesserts?item='+data.item)
       .then((response) => {
         console.log(response.data);

       })
      .catch((error)=>{
         console.log(error);
      });

        this.incrementItemsCart();
};

addSmoothieToCart = (data) => {
        // Backend Api to feed into the database

        console.log("Inside addSmoothieToCart", data);

        axios.get('http://52.52.176.4:5000/addToCartDesserts?item='+data.item)
       .then((response) => {
         console.log(response.data);

       })
      .catch((error)=>{
         console.log(error);
      });

        this.incrementItemsCart();
};

addTeaToCart = (data) => {
        // Backend Api to feed into the database

        console.log("Inside addTeaToCart", data);

        axios.get('http://52.52.176.4:5000/addToCartDesserts?item='+data.item)
       .then((response) => {
         console.log(response.data);

       })
      .catch((error)=>{
         console.log(error);
      });

        this.incrementItemsCart();
};

redirectToMainPage = () => {
  this.props.history.push('/mainPage');
}


    render() {
        return (
            <div>

                <Route exact path="/" render={() => (
                    <div>
                        <Home redirectToMainPage={this.redirectToMainPage}/>
                    </div>
                )}/>

                <Route exact path="/mainPage" render={() => (
                    <div>
                        <MainPage checkCart={this.checkCart} itemsCount={this.state.totalItems} addCoffeeToCart={this.addCoffeeToCart} addTeaToCart={this.addTeaToCart} addSmoothieToCart={this.addSmoothieToCart} addBreadToCart={this.addBreadToCart} addDessertToCart={this.addDessertToCart} findStores={this.findStores} incrementLikesCoffee={this.incrementLikesCoffee} incrementLikesTea={this.incrementLikesTea} incrementLikesBread={this.incrementLikesBread} incrementLikesDesserts={this.incrementLikesDesserts}/>
                    </div>
                )}/>

                <Route exact path="/Checkout" render={() => (
                    <div>

                    </div>
                )}/>

                <Route exact path="/cart" render={() => (
                    <div>
                      <Cart processOrder={this.processOrder}/>
                    </div>
                )}/>

                <Route exact path="/stores" render={() => (
                    <div>
                      <Stores findStores={this.findStores} stores={this.state.stores}/>
                    </div>
                )}/>

            </div>
        );
    }
}

export default withRouter(Routing);
