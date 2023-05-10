import React from 'react';
import { Fragment } from "react";
import { BrowserRouter, Route, Routes} from "react-router-dom";
import ListagemProdutosCadastrados from '../componentes/ListagemProdutosCadastrados';
import SellingPhones from '../views/SellingPhones';
import CadastrarProduto from '../views/cadastrarProduto';
import PhonesToSell from '../views/PhonesToSell';
import ViewDoCarrinho from '../views/ViewDoCarrinho';
import LoginPage from '../componentes/LoginPage';

// import { Container } from './styles';

function RoutesApp(props) {
  return (
    <BrowserRouter>
    <Fragment>
        <Routes>
            <Route exact path="/listarProdutosCadastraddos" element = {<SellingPhones></SellingPhones>}></Route>
            <Route exact path = "/cadastarProduto" element = {<CadastrarProduto></CadastrarProduto>} />
            <Route exact path = "/vendaDosCelures" element = {<PhonesToSell></PhonesToSell>}></Route>
            <Route exact path = "/carrinho" element = {<ViewDoCarrinho></ViewDoCarrinho>}></Route>
            <Route path="/" element = {<LoginPage></LoginPage>} />
            <Route path="*" element = {<LoginPage></LoginPage>} />
        </Routes>
    </Fragment>
    </BrowserRouter>
    );

}

export default RoutesApp;