import React from 'react';
import '../css/PhoneList.css';
import Phone from './phone';

function PhoneList() {
  const phones = [
    { id: 1, name: 'Samsung Galaxy S21', price: 4999 },
    { id: 2, name: 'iPhone 12 Pro', price: 7999 },
    { id: 3, name: 'Xiaomi Mi 11', price: 2999 },
    { id: 4, name: 'Google Pixel 5', price: 4499 },
  ];

  return (
    <div className="phone-list-page">
      <h1>Phones for Sale</h1>
      <div>
        {phones.map((phone) => (
            <Phone id = {phone.id} name = {phone.nome} preco = {phone.price} quantidade = {3} ></Phone>
            ))}
      </div>
    </div>
  );
}

export default PhoneList;

