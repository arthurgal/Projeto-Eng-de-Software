import React from 'react'
import Sidebar from '../componentes/sidebar'
import '../css/sellingPhones.css'
import ListagemProdutosCadastrados from '../componentes/ListagemProdutosCadastrados'

export default function SellingPhones() {
  return (
    <div className="div-padrao">
      <Sidebar usuario = "vendedor"></Sidebar>
      <div className='div-main'>
         <div className='div-title'>
            <h1>Produtos cadastrados</h1>
            <a href="/cadastarProduto" className='button-novo-produto' >novo Produto</a>
         </div>
         <div className='div-listing'>
            <ListagemProdutosCadastrados nome = "nome" valor = "valor"></ListagemProdutosCadastrados>
            <ListagemProdutosCadastrados nome = "nome" valor = "valor"></ListagemProdutosCadastrados>
         </div>
      </div>
    </div>
  )
}
