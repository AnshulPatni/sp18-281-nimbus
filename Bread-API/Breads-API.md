# Breads API 

1. **Title** : inventoryHandlerBreads
Description : API retrieves bread menu items from database
URL :  /inventoryBreads
Method : GET Response Codes : 200(OK), 401(Webpage doesn’t exist), 500(Internal server error)
 
2. **Title** : searchInventoryBreads
Description : API retrieves menu item searched by user
URL : /searchInventoryBreads
Method : GET
Response Codes : 200(OK), 401(Webpage doesn’t exist), 500(Internal server error)

3. **Title** : starbucksAddToCartHandlerBreads
I.	Description : API adds selected items to cart
II.	URL : /addToCartBreads
III.	Method : PUT
IV.	Response Codes : 200(OK), 401(Webpage doesn’t exist), 500(Internal server error)

4. **Title** : cartHandlerBreads
I.	Description : API retrieves items  that have been added to cart
II.	URL : /cartItemsBreads
III.	Method : GET
IV.	Response Codes : 200(OK), 401(Webpage doesn’t exist), 500(Internal server error)