Get Item
Method: GET
Path: /items/{id}
Description: Retrieves an item by its ID.
Parameters:
id (path parameter): The ID of the item to retrieve.
Responses:
200 OK: Returns the item details.
404 Not Found: Returns an error if the item is not found.

Add Item
Method: POST
Path: /items
Description: Adds a new item to the inventory.
Request Body:
JSON
{
"name": "Product A",
"quantity": 100,
"warehouse_id": "WH1"
}
Responses:
201 Created: Returns the newly created item.
400 Bad Request: Returns an error if the request body is invalid.
500 Internal Server Error: Returns an error if there's an issue adding the item.

Update Item
Method: PUT
Path: /items/{id}
Description: Updates an existing item in the inventory.
Parameters:
id (path parameter): The ID of the item to update.
Request Body:
JSON
{
"name": "Updated Product A",
"quantity": 150
}

Delete Item
Method: DELETE
Path: /items/{id}
Description: Deletes an item from the inventory.
Parameters:
id (path parameter): The ID of the item to delete.
Responses:
200 OK: Returns a success message.
404 Not Found: Returns an error if the item is not found.
500 Internal Server Error: Returns an error if there's an issue deleting the item.

Get Inventory by Warehouse
Method: GET
Path: /inventory/{warehouse_id}
Description: Retrieves a list of items in a specific warehouse.
Parameters:
warehouse_id (path parameter): The ID of the warehouse.
Responses:
200 OK: Returns a list of items.
500 Internal Server Error: Returns an error if there's an issue retrieving the inventory.