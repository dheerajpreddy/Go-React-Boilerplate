import React, { Component } from 'react';
import './ViewPeople.css';

class ViewPeople extends Component {
  constructor() {
    super();
    this.state = {
      data: []
    }
  }

  componentDidMount() {
    // fetch('http://localhost:8080/people/', {
    //   method: 'GET',
    //   mode: "no-cors",
    // })
      // .then(results => {
      //   console.log(results);
        // console.log(d);
        // return results.json();
      // })
    //   .then( data => {
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
    // console.log(this.state.allUsers);
    const request = new Request('http://127.0.0.1:8080/people/');
    fetch(request)
      .then((response, err) => {
        console.log(response);
        return response.json();
      })
        .then(data => this.setState({data: data}));

    // fetch(request)
    //   .then((response, err) => console.log(response))
    //     .then(data => this.setState({data: data}));
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">View All People</h1>
        </header>

        <table className="table">
          <tbody>{this.state.data.map(function(item, key) {
               return (
                  <tr key = {key}>
                      <td>{item.id}</td>
                      <td>{item.firstname}</td>
                      <td>{item.lastname}</td>
                      <td>{item.city}</td>
                  </tr>
                )

             })}</tbody>
       </table>
      </div>
    );
  }
}

export default ViewPeople;
