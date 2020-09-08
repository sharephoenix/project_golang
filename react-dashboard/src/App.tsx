import React from 'react';
import logo, { ReactComponent } from './logo.svg';
import './App.css';
import { tsThisType } from '@babel/types';

interface IAppComState{
  nicknames: string
}

interface IAppProps { 
  address?: string
}

class TTest extends React.Component<IAppProps,IAppComState> { 

  nickname: string = 'alexluan'
  address?: string = '--'

  constructor(props: IAppProps) { 
    super(props)
    this.address = props.address
    this.nickname = 'hhhhhh'
    this.state = {
      nicknames: this.props.address || 'vvvvvv'
    }
    this.clickHandler = this.clickHandler.bind(this);
  }

  onStart() { 
    console.log('ssstart')
  }

  test(): string { 
    return 'this is my custom function!!!'
  }

  clickHandler() { 
    console.log('clickHandler', this.setState)
    this.setState({
      nicknames: "ericjump"
    })
    // this.nickname = 'lllllllll'
  }

  render() { 
    return <div className="App">
      {this.test()}
      {this.state.nicknames}
      {this.props.address}
    <header className="App-header">
      <img src={logo} className="App-logo" alt="logo" />
        <p onClick={this.clickHandler}>
        Edit <code>src/App.tsx</code> and save to reload.
      </p>
      <a
        className="App-link"
        href="https://reactjs.org"
        target="_blank"
        rel="noopener noreferrer"
      >
        
        Learn React
      </a>
    </header>
  </div>
  }
}

export default TTest;
