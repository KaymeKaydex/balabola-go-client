package balabola

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctx := context.TODO()

	c := New(ctx)

	nonsense, err := c.TalkNonsense("Акции bitop выросли на 5%", TextStyleWithoutStyle, 1)
	if err != nil {
		log.Default().Print(err)
	}

	fmt.Println(nonsense)

	assert.NotNil(t, c)
}
