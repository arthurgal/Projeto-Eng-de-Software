import React from 'react'
import "../css/sidebar.css"
export default function Sidebar(props) {
  return (
    <div className='sideBar'>
        { props.usuario == "vendedor" ?
        
            <div className='menu'>
                  <a className='itemMenu'>Meus produtos</a>     
                  <a className='itemMenu'>Produtos vendidos</a>     
            </div> 
        
        : ""
        }

        {
             props.usuario == "comprador" ?
             <div className='menu'>
                     
                     <a className='itemMenu'>Produtos a serem vendidos</a>   
                     
             </div> 
             : ""
        }
    </div>
  )
}
