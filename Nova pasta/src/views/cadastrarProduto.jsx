import React from 'react'
import Sidebar from '../componentes/sidebar'
import '../css/sellingPhones.css'
import ListagemProdutosCadastrados from '../componentes/ListagemProdutosCadastrados'
import NovoProduto from '../componentes/NovoProduto'

export default function CadastrarProduto() {
  return (
    <div className="div-padrao">
      <Sidebar usuario = "vendedor"></Sidebar>
      <div className='div-main'>
         <div className='div-title'>
            <h1>Cadastrar Produto</h1>
          { //<a href="" className='button-novo-produto' >Salvar Produto</a>
            }
         </div>
         <div className='div-listing'>
            <NovoProduto></NovoProduto>
         </div>
      </div>
    </div>
  )
}