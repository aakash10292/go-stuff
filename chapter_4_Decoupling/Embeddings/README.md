## Embedding

Embedding types provide the final piece of sharing and reusing state and behavior between types. Through the use of inner type promotion, an inner type's fields and methods can be directly accessed by references of the outer type.

## Notes

* Embedding types allow us to share state or behavior between types.
* The inner type never loses its identity.
* This is not inheritance.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* The outer type can override the inner type's behavior.
