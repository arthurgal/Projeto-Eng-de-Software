import React from 'react'
import imgEditar from '../assets/icons8-editar-50.png'
import imgExcluir from '../assets/icons8-remover-48.png'
import '../css/listingprodutoscadastrados.css'

export default function ListagemProdutosCadastrados(props) {
  return (
    <div className='div-listing-produtos'>
      <div>
          {props.nome}
      </div>
      <div>
        {props.valor}
      </div>
       <div className='buttons-div'> 
          <img className='img' src={imgEditar} alt="editar" />
          <img className='img' src={imgExcluir} alt = "excluir" />
       </div>
    </div>
  )
}
