import React, { useState } from "react";
import Beverage from "../interfaces/beverage_interface";

import "./addproduct.css";

const AddProduct = (props: any) => {
  const [newBeverage, setNewBeverage] = useState<Beverage>({
    type: "Coffee",
    roast_degree: 1,
  } as Beverage);

  // Handle input change
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewBeverage({ ...newBeverage, [e.target.name]: e.target.value });
  };

  // Handle select change
  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setNewBeverage({ ...newBeverage, [e.target.name]: e.target.value });
  };

  // Handle form submit
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    props.addBeverage(newBeverage);
  };

  return (
    <form onSubmit={handleSubmit} className="addbeverage-form">
      <div className="addbeverage-form-header">
        <h2>Lisää uusi tuote</h2>
      </div>
      <div className="addbeverage-form-body">
        <div className="addbeverage-form-left">
          <input
            className="addbeverage-form-input"
            type="text"
            name="name"
            placeholder="Nimi"
            onChange={handleInputChange}
          />
          <input
            className="addbeverage-form-input"
            type="number"
            name="weight"
            step={0.01}
            placeholder="Paino (g)"
            onChange={handleInputChange}
          />
          <input
            className="addbeverage-form-input"
            type="number"
            name="price"
            step={0.01}
            placeholder="Hinta (€)"
            onChange={handleInputChange}
          />
        </div>
        <div className="addbeverage-form-right">
          <label htmlFor="roast_degree">Paahtoaste</label>
          <select name="roast_degree" onChange={handleSelectChange}>
            <option value="1">1</option>
            <option value="2">2</option>
            <option value="3">3</option>
            <option value="4">4</option>
            <option value="5">5</option>
          </select>
          <label htmlFor="type">Tyyppi</label>
          <select name="type" onChange={handleSelectChange}>
            <option value="Coffee">Coffee</option>
            <option value="Tea">Tea</option>
          </select>
          <button type="submit">Lisää</button>
        </div>
      </div>
    </form>
  );
};

export default AddProduct;
