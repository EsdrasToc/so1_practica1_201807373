import 'bootstrap/dist/css/bootstrap.css'
import React from 'react'
import './App.css'
import Table from './Table'

class App extends React.Component{

    constructor(props){

        super(props)
        
        this.state = {
            placa:"",
            marca:"",
            modelo:0,
            serie:"",
            color:"",
            filter: 0,
            cars: []
        };

        this.path = "http://localhost:3030"

        fetch(this.path+'/hola',{
            method: 'POST', 
            mode: 'cors', 
            body: JSON.stringify({})
        })

        this.handleInputChange = this.handleInputChange.bind(this);
        this.Create = this.Create.bind(this);
    }

    Create(){
        if(this.state.placa === "" || this.state.marca === "" || this.state.modelo === 0 || this.state.serie === "" || this.state.color === ""){
            alert("datos incompletos");
        }else{

            const newCar = {
                placa:this.state.placa,
                marca: this.state.marca,
                modelo:this.state.modelo,
                serie:this.state.serie,
                color:this.state.color,
            }

            fetch(this.path+'/create',{
                method: 'POST', 
                mode: 'cors', 
                body: JSON.stringify(newCar)
            });
        
            console.log("lalalal")
        }
    }

    

    handleInputChange(event){
        const target = event.target;
        const name = target.name;
        const value = target.value;
        
        if(name === "modelo"){
            this.setState({
                [name]:parseInt(value)
            });
        }else{
            this.setState({
                [name]: value
            });
        }
    }

    render(){
        return(
            <div>
            <header className="p-3 text-bg-dark">
                <div className="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
                    <a href="/" className="d-flex align-items-center mb-2 mb-lg-0 text-white text-decoration-none">
                    </a>

                    <div className='form'>
                        <ul className="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
                        <li className="createElement">Placa : <input type="text" name='placa' className="new" value={this.state.placa} onChange={this.handleInputChange}/></li>
                        <li className="createElement">Marca : <input type="text" name='marca' className="new" value={this.state.marca} onChange={this.handleInputChange}/></li>
                        <li className="createElement">Modelo : <input type="number" name='modelo' className="new" value={Number(this.state.modelo)} onChange={this.handleInputChange}/></li>
                        <li className="createElement">Serie : <input type="text" name='serie' className="new" value={this.state.serie} onChange={this.handleInputChange}/></li>
                        <li className="createElement">Color : <input type="text" name='color' className="new" value={this.state.color} onChange={this.handleInputChange}/></li>
                        <li className="createElement"><button className="btn btn-primary" onClick={this.Create}>Create</button></li>
                        </ul>
                    </div>
                </div>
            </header>
                <Table />
            </div>
        )
    }
}

export default App;