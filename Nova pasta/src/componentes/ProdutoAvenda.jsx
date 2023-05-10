import React from 'react';
import imgCelular from "../assets/celular.png" 
// import { Container } from './styles';
import "../css/phone.css"
function ProdutoAvenda(props) {
  return <div className='div-bloco'>
        <div>
          <img src={imgCelular} alt="" />
        </div>
        <div>
            <h2>{props.nome}</h2>
        </div>
        <div>
            <h2>{props.descricao}</h2>
        </div>
        <div>
            <h2> R$ {props.preco}</h2>
        </div>
        <div>
           {props.jaAdd != undefined ? "" : <button>adcionar ao carrinho</button>  }
           
        </div>
  </div>;
}

export default ProdutoAvenda;