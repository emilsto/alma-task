import Beverage from '../interfaces/beverage_interface';

interface Props {
    beverage: Beverage;
}

const ProductCard = ({ beverage }: Props) => {
    return (
        <tr>
            <td>{beverage.name}</td>
            <td>{beverage.weight}</td>
            <td>{beverage.price}</td>
            <td>{beverage.roast_degree}</td>
            <td>{beverage.type === 'Coffee' ? 'Kahvi' : 'Tee'
        }</td>
        </tr>
    );
};

export default ProductCard;