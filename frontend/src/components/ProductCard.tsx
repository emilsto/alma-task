import Beverage from '../interfaces/beverage_interface';

interface Props {
    beverage: Beverage;
}

const ProductCard = ({ beverage }: Props) => {
    return (
        <tr>
            <td>{beverage.name}</td>
            <td>{beverage.weight_grams}</td>
            <td>{beverage.price}</td>
            <td>{beverage.roast_degree}</td>
            <td>{beverage.type}</td>
        </tr>
    );
};

export default ProductCard;