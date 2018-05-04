import React, {Component} from 'react';
import './Home.css';


class Home extends Component {

  constructor(props) {
      super(props);
      this.state={
        pageView: "coffee",
        dummyData: "lelo mc"
      }


    }



render() {
  return (
    <div className="homepage">
        <div className="inner animated fadeIn" id="fade">
          <h3> Welcome To Starbucks Online App.</h3>
          <br/>
          <br/>
          <button className="btn btn-default butt" id="fade1" onClick={this.props.redirectToMainPage}><div className="buttText">Order Online</div></button>
        </div>
    </div>
);
}
  }


export default Home;
