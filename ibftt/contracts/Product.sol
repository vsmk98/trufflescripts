pragma solidity ^0.4.18;

contract Product {
  bytes32 public name;
  bytes12 public productId;

  function Product (bytes32 _name, bytes12 _productId) public {
    name = _name;
    productId = _productId;
  }

  function getInfo () external view returns (bytes32, bytes12) {
    return (name, productId);
  }
}
