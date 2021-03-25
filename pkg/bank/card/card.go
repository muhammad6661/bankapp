package card

import (
	"bank/pkg/bank/types"
)

func Total (card [] types.Card) types.Money {
    sum:=0
     for _,operation :=range card {
       if(operation.Active&&operation.Balance>0){
        sum+=int(operation.Balance)
       }
     }
  return  types.Money(sum)
    }