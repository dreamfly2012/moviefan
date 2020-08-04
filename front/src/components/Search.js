import React, { Component } from 'react'
import './search.css'
import {getMovie} from '../api/search.js'

export default class Search extends Component {
    constructor(props){
        super(props)
        this.state={
            movie:{}
        }
    }
    componentDidMount(){
        getMovie({'params':{'name':'战狼2'}}).then(res=>{
            this.setState({
                movie:res.info
            })
        }).catch(function (err) {
            console.log(err);
        });
        
    }
    
    render() {
        return (
            <div>
                <input type="search" placeholder="影片搜索" />
                <p>{this.state.movie.Summary}</p>
            </div>
        )
    }
}
