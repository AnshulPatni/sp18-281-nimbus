package main

type gumballMachine struct {
	item        string 	
	count   	int    	
	cost 		int	    
	status 		int	
}

type order struct {
	Id             	string 	
	OrderStatus 	string	
}

var orders map[string] order
