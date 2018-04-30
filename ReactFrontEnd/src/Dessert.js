import React, {Component} from 'react';
import './itempage.css';

class Dessert extends Component {

constructor(props) {
      super(props);
this.state={
  desserts: [
    {
      cost: 5.2,
      item: "indian chai",
      likes: 0,
      status: 0,
      typeService: "tea"
    },
    {
      cost: 4.0,
      item: "american chai",
      likes: 0,
      status: 0,
      typeService: "tea"
    }
  ]
}
}

componentWillMount() {

console.log("lets check",this.props.desserts);
this.setState({
  desserts: this.props.desserts
})

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

{this.props.desserts.map(dessert => (
          <div className="row coffeerow" style={{marginTop: 3}}>
              <div className="col-lg-2 desc">
                {dessert.likes}&nbsp; <i className="fas fa-heart fa-1x" style={{color: "red", fontSize: 18}} onClick={() => {this.props.incrementLikesDesserts({item: dessert.item, likes: dessert.likes})}}></i>
              </div>
              <div className="col-lg-4 desc">
                {dessert.item}
              </div>
              <div className="col-lg-3 desc">
                {dessert.cost}
              </div>
              <div className="col-lg-3 desc" style={{paddingTop: 5}}>
                <button type="button" className="btn btn-primary" onClick={()=>{this.props.addDessertToCart({item: dessert.item})}}>Add to Cart </button>
              </div>
          </div>
))}


        </div>
        <br/>

    </div>
);
}
  }


export default Dessert;
