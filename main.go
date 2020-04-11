package pub

import (
	"context"
	"github.com/Snowlights/pub/adapter"
)

func main(){
	ctx := context.Background()
	adapter.Prepare(ctx)
}

