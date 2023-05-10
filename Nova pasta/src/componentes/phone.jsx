import React from 'react';
import "../css/phone.css"
// import { Container } from './styles';

function Phone(props) {
  return(
    <div className='conteinerProduto'>
        <div>
            {props.nome}
        </div>
        <div>
            {props.preço}
            <div>
                {props.quantidade}    
            </div>    
        </div>
        
  </div>);
}

export default Phone;