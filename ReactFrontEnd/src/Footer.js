import React, {Component} from 'react';
import './Footer.css';

import FooterLogo from './footer1.png';

class Footer extends Component {

state={

}

render() {
  return (
    <div className="footerPage">
      <div className="container-fluid">
          <div className="row">
              <div className="coll-lg-6">
                <img src={FooterLogo} alt="" />
              </div>
              <div className="coll-lg-4">
              </div>
              <div className="coll-lg-2">
              </div>
          </div>
      </div>
    </  div>
);
}
}


export default Footer;
