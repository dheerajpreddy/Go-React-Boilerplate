import React, { Component } from 'react';
import './ViewPeople.css';

class ViewPeople extends Component {
  constructor() {
    super();
    this.state = {
      allUsers: {},
    }
  }

  componentDidMount() {
    // fetch('http://localhost:8080/people/')
    //   .then(results => {
    //     return results.json();
    //   }).then( data => {
    //     for(var x in data) {
    //       console.log(x);
    //       let allUsers = [];
    //       allUsers.push(
    //         <tr>
    //           <td>{x.id}</td>
    //           <td>{x.firstname}</td>
    //           <td>{x.lastname}</td>
    //           <td>{x.city}</td>
    //         </tr>
    //       )
    //       this.setState(allUsers: allUsers);
    //     }
    //   })
    fetch('http://localhost:8080/people/')
      .then(results => {
        console.log(results);});
    // console.log(this.state.allUsers);
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">View All People</h1>
        </header>
    
      </div>
    );
  }
}

export default ViewPeople;
