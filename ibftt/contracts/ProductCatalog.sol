pragma solidity ^0.4.18;

import './Product.sol';

contract ProductCatalog {
  Product[] public products;
  mapping(bytes12 => Product) public productIdToProduct;

  event ProductCreated(address productAddress, bytes32 name, bytes32 productId);

  function createProduct(bytes32 name, bytes12 productId) external returns (address) {
    Product product = new Product(name, productId);
    products.push(product);
    productIdToProduct[productId] = product;

    ProductCreated(product, name, productId);
  }

  function getProductCount () external view returns (uint) {
    return products.length;
  }
}
