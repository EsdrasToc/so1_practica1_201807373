import 'bootstrap/dist/css/bootstrap.css'
import React from 'react'


class Row extends React.Component{
    constructor(props){
        super(props)

        this.state = {
            car : props.car
        };
    }
    render(){
        return(
            <tr>
                <td>{this.state.car.placa}</td>
                <td>{this.state.car.marca}</td>
                <td>{this.state.car.modelo}</td>
                <td>{this.state.car.serie}</td>
                <td>{this.state.car.color}</td>
                <td><button type="button" className="btn btn-warning">Delete</button></td>
                <td><button type="button" className="btn btn-primary">Update</button></td>
            </tr>
        );
    }
}

class Table extends React.Component{
    constructor(props){
        super(props)

        this.state = {
            cars: [],
            filter: 0,
            filterValue: ""
        };
        
        this.handleFilterChange = this.handleFilterChange.bind(this);
        this.handleFilterValue = this.handleFilterValue.bind(this);
        this.ReadAll = this.ReadAll.bind(this);
        this.Read = this.Read.bind(this);
    }

    async ReadAll(){
        const response = await fetch('http://localhost:3030/readall');
        const data = await response.json();
        
        this.setState({
            cars: data
        });
    }

    async Read(){
        if(this.state.filter === 0){
            this.ReadAll();
            return;
        }

        const response = await fetch('http://localhost:3030/filter/'+this.state.filter+"/"+this.state.filterValue)
        const data = await response.json();
        this.setState({
            cars:data
        });
    }

    handleFilterChange(event){
        const name = event.target.name;
        const value = event.target.value;
        this.setState(
            {
                [name]:parseInt(value)
            }
        );
    }

    handleFilterValue(event){
        const name = event.target.name;
        const value = event.target.value;

        this.setState(
            {
                [name]:value
            }
        );
    }

    render(){
        const cars = this.state.cars.map((c)=>
            <Row car={c} key={c.placa}/>
        );
        return(
            <div className="container">
                <div className='main-row'>
                    <form name="read">
                    <ul className="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
                        <li className="createElement">Filtrar : <input type="text" name='filterValue' className="new" onChange={this.handleFilterValue}/></li>
                        <li className="createElement">por : <select name='filter' className="new" value={this.filter} onChange={this.handleFilterChange}>
                                <option value="0">Ninguno</option>
                                <option value="1">Marca</option>
                                <option value="2">Modelo</option>
                                <option value="3">Color</option>
                            </select></li>
                        <li className="createElement"><button type="button" className="btn btn-success" onClick={this.Read}>Filter</button></li>
                        </ul>
                    </form>
                </div>
                <div className="row">
                    <table className="table">
                        <thead>
                            <tr>
                            <th scope="col">Placa</th>
                            <th scope="col">Marca</th>
                            <th scope="col">Modelo</th>
                            <th scope="col">Serie</th>
                            <th scope="col">Color</th>
                            </tr>
                        </thead>
                        <tbody>
                            {cars}
                        </tbody>
                    </table>
                </div>
            </div>
            
        );
    }
}

export default Table;