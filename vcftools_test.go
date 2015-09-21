package vcftools

import (
"testing"
"github.com/stretchr/testify/assert"
)

func TestVariantInLine(t *testing.T) {
    arr := []string {"./.:26,0:26:69.22:0,69,898",
                     "0/1:26,0:26:69.22:0,69,898"}
    assert.Equal(t, true, VariantInLine(arr))

    arr = []string {"./.:26,0:26:69.22:0,69,898",
                    "0/0:26,0:26:69.22:0,69,898"}
    assert.Equal(t, false, VariantInLine(arr))
}


