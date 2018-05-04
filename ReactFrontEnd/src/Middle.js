import React, {Component} from 'react';
import './mainpage.css';
import Coffee from './Coffee';
import Tea from './Tea';
import Bread from './Bread';
import Dessert from './Dessert';
import Smoothie from './Smoothie';

class Middle extends Component {

state={

}

componentWillMount() {
  console.log("lets check it in Middle", this.props.coffees)
}

render() {
  if(this.props.pageView === "desserts")
  {
      return (
        <div className="content1">
          <Dessert desserts={this.props.desserts} addDessertToCart={this.props.addDessertToCart} incrementLikesDesserts={this.props.incrementLikesDesserts}/>
        </div>
    );
  }
  else if(this.props.pageView === "tea"){
    return (
      <div className="content1">
        <Tea teas={this.props.teas} addTeaToCart={this.props.addTeaToCart} incrementLikesTea={this.props.incrementLikesTea}/>
      </div>
  );
  }
  else if(this.props.pageView === "breads") {
    return (
      <div className="content1">
        <Bread breads={this.props.breads} addBreadToCart={this.props.addBreadToCart} incrementLikesBread={this.props.incrementLikesBread}/>
      </div>
  );
  }
  else if(this.props.pageView === "snacks") {
    return (
      <div className="content1">
        <Smoothie smoothies={this.props.smoothies} addSmoothieToCart={this.props.addSmoothieToCart} incrementLikesSmoothies={this.props.incrementLikesSmoothies}/>
      </div>
  );
  }
  else {
    return (
      <div className="content1">
          <Coffee coffees={this.props.coffees} addCoffeeToCart={this.props.addCoffeeToCart} incrementLikesCoffee={this.props.incrementLikesCoffee}/>
      </div>
  );
  }
}
}


export default Middle;
