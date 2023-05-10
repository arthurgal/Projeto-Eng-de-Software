import React from 'react'
import Sidebar from '../componentes/sidebar'
import '../css/sellingPhones.css'
import ListagemProdutosCadastrados from '../componentes/ListagemProdutosCadastrados'
import ProdutoAvenda from '../componentes/ProdutoAvenda'

export default function ViewDoCarrinho() {
  return (
    <div className="div-padrao">
      <Sidebar usuario = "comprador"></Sidebar>
      <div className='div-main'>
         <div className='div-title'>
            <h1>Produtos cadastrados</h1>
            <a href="/cadastarProduto" className='button-novo-produto' >novo Produto</a>
         </div>
         <div className='div-listing'>
                <ProdutoAvenda id= {3} nome = "Sansung" preco = "39999" descricao = "celular muito bom" jaAdd = {true}></ProdutoAvenda>
         </div>
      </div>
    </div>
  )
}
