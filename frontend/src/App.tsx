import { useState, useEffect } from 'react';
import ProductCard from './components/ProductCard';
import './App.css';

// Interface
import Beverage from './interfaces/beverage_interface';

const App = () => {
  const [beverages, setBeverages] = useState<Beverage[]>([]);

  useEffect(() => {
    const getBeverages = async () => {
      const response = await fetch('http://localhost:5000/api/v1/beverages');
      const data = await response.json();
      setBeverages(data);
    };
    getBeverages();

  }, []);


  return (
    <div className="App">
      <header className="App-header">
      <table>
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
          {beverages.map((beverage) => {
            return <ProductCard beverage={beverage} key={beverage.id} />;
          })
          }
          </tbody>
            </table>
      </header>
    </div>
  );
}

export default App;
