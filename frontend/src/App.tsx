import { useState, useEffect } from "react";
import ProductCard from "./components/ProductCard";
import "./App.css";

import AddProduct from "./components/AddProduct";

// Interface
import Beverage from "./interfaces/beverage_interface";


const App = () => {
  const [beverages, setBeverages] = useState<Beverage[]>([]);
  const [search, setSearch] = useState<string>("");

  useEffect(() => {
    const getBeverages = async () => {
      try {
        const response = await fetch(
          `http://localhost:5001/api/v1/beverages${search ? `/${search}` : ""}`
        );
        const data = await response.json();
        setBeverages(data);
      } catch (error) {
        // Handle error
        console.log(error);
      }
    };
    getBeverages();
  }, [search]);

  const addBeverage = (beverage: Beverage) => {
    const postBeverage = async () => {
      const response = await fetch("http://localhost:5001/api/v1/beverages", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        // Format the data to JSON, make sure to use correct types
        body: JSON.stringify({
          name: beverage.name,
          weight: Number(beverage.weight),
          price: Number(beverage.price),
          roast_degree: Number(beverage.roast_degree),
          type: beverage.type,
        }),
      });
      const data = await response.json();
      setBeverages([...beverages, data]);
    };
    postBeverage();
  };

  return (
    <div className="App">
      <h1>Kahvi- ja teetietopankki</h1>

      <header className="main-view">
        <div className="add-product">
          <AddProduct addBeverage={addBeverage} />
        </div>
        <div className="search-and-view">
          <input
            type="text"
            placeholder="Hae tuotteita..."
            className="search-products"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
          <div className="table-holder">
            <table className="styled-table">
              <thead>
                <tr>
                  <th>Nimi</th>
                  <th>Paino (g)</th>
                  <th>Hinta (€)</th>
                  <th>Paahtoaste</th>
                  <th>Tyyppi</th>
                </tr>
              </thead>
              <tbody>
                {beverages &&
                  beverages.map((beverage) => {
                    return (
                      <ProductCard beverage={beverage} key={beverage.id} />
                    );
                  })}
              </tbody>
            </table>
          </div>
        </div>
      </header>
    </div>
  );
};

export default App;
