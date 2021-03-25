package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleTotal(){
	card:=[]types.Card{
		{
		Balance: 10_000_00,
		Active: true,
	   },
	   {
		   Balance: 10_000_00,
          Active: true,
		},
}
  fmt.Println(Total(card))
  //Output: 2000000
}
