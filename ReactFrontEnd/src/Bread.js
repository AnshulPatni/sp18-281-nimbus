import React, {Component} from 'react';
import './itempage.css';

class Bread extends Component {

constructor(props) {
      super(props);
this.state={
  breads: [
    {
      cost: 5.2,
      item: "indian bread",
      likes: 0,
      status: 0,
      typeService: "tea"
    },
    {
      cost: 4.0,
      item: "american bread",
      likes: 0,
      status: 0,
      typeService: "tea"
    }
  ]
}
}

componentWillMount() {

console.log("lets check breads: ",this.props.breads);
this.setState({
  breads: this.props.breads
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

{this.props.breads.map(bread => (
          <div className="row coffeerow" style={{marginTop: 3}}>
              <div className="col-lg-2 desc">
                {bread.likes}&nbsp; <i className="fas fa-heart fa-1x" style={{color: "red", fontSize: 18}} onClick={() => {this.props.incrementLikesBread({item: bread.item, likes: bread.likes})}}></i>
              </div>
              <div className="col-lg-4 desc">
                {bread.item}
              </div>
              <div className="col-lg-3 desc">
                {bread.cost}
              </div>
              <div className="col-lg-3 desc" style={{paddingTop: 5}}>
                <button type="button" className="btn btn-primary" onClick={()=>{this.props.addBreadToCart({item: bread.item})}}>Add to Cart </button>
              </div>
          </div>
))}


        </div>
        <br/>

    </div>
);
}
  }


export default Bread;
