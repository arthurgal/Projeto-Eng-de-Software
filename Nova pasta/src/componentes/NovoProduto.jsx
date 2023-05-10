import React, { useState } from 'react';

// import { Container } from './styles';

function NovoProduto() {
    const [nome,setNome] = useState("")
    const [valor,setValor] = useState("")
    const [Descricao,setDescricai] = useState("")
    const [qtd,setQtd] = useState("")
  return(
  <div>
    <form action="" method="post">
        <label>
            Nome:
            <input type="text" name="name" value={nome} onChange={(e)=> setNome(e.target.value) } />
        </label>
        <label>
            Valor:
            <input type="text" name="name" value={valor} onChange={(e)=> setValor(e.target.value) }  />
        </label>
        <label>
            Descricao:
            <input type="text" name="name"  value={Descricao} onChange={(e)=> setDescricai(e.target.value) }/>
        </label>
        <label>
            Qtd:
            <input type="text" name="name" value={qtd} onChange={(e)=> setQtd(e.target.value) } />
        </label>
        
        <input type="submit" value="Enviar" />

    </form>
  </div>)
}

export default NovoProduto;