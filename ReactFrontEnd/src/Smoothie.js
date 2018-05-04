import React, {Component} from 'react';
import './itempage.css';

class Smoothie extends Component {

constructor(props) {
      super(props);
this.state={
  smoothies: [
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

console.log("lets check smoothies",this.props.smoothies);
this.setState({
  smoothies: this.props.smoothies
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

{this.props.smoothies.map(smoothie => (
          <div className="row coffeerow" style={{marginTop: 3}}>
              <div className="col-lg-2 desc">
                {smoothie.likes}&nbsp; <i className="fas fa-heart fa-1x" style={{color: "red", fontSize: 18}} onClick={() => {this.props.incrementLikesSmoothies({item: smoothie.item, likes: smoothie.likes})}}></i>
              </div>
              <div className="col-lg-4 desc">
                {smoothie.item}
              </div>
              <div className="col-lg-3 desc">
                {smoothie.cost}
              </div>
              <div className="col-lg-3 desc" style={{paddingTop: 5}}>
                <button type="button" className="btn btn-primary" onClick={()=>{this.props.addSmoothieToCart({item: smoothie.item})}}>Add to Cart </button>
              </div>
          </div>
))}


        </div>
        <br/>

    </div>
);
}
  }


export default Smoothie;
