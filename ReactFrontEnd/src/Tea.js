import React, {Component} from 'react';
import './itempage.css';

class Tea extends Component {

constructor(props) {
      super(props);
this.state={
  teas: [
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

console.log("lets check",this.props.teas);
this.setState({
  teas: this.props.teas
})

}

addTeaToCart(data) {

  console.log("Inside AddTeaToCart");

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

{this.props.teas.map(tea => (
          <div className="row coffeerow" style={{marginTop: 3}}>
              <div className="col-lg-2 desc">
                {tea.likes}&nbsp; <i className="fas fa-heart fa-1x" style={{color: "red", fontSize: 18}} onClick={() => {this.props.incrementLikes({item: tea.item, likes: tea.likes})}}></i>
              </div>
              <div className="col-lg-4 desc">
                {tea.item}
              </div>
              <div className="col-lg-3 desc">
                {tea.cost}
              </div>
              <div className="col-lg-3 desc" style={{paddingTop: 5}}>
                <button type="button" className="btn btn-primary" onClick={()=>{this.props.addTeaToCart({item: tea.item})}}>Add to Cart </button>
              </div>
          </div>
))}


        </div>
        <br/>

    </div>
);
}
  }


export default Tea;
