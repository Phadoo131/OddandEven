package OddandEven

import (
	"testing"

	"github.com/Phadoo131/OddandEven/util"
	"github.com/stretchr/testify/require"
)

func TestEven (t *testing.T){
	randNum := util.RandomSliceOfInt(10)

	test := Even(randNum)
	for _, i := range test{
		require.Equal(t, i%2, 0)
	}
}

func TestOdd (t *testing.T){
	randNum := util.RandomSliceOfInt(10)

	test := Odd(randNum)
	for _, i := range test{
		require.Equal(t, i%2, 1)
	}
}

func TestCalculate(t *testing.T){
	randNum := util.RandomSliceOfInt(10)
	ans := 0

	test := Calculate(randNum)

	for _, i := range randNum{
		ans += i
	}

	require.Equal(t, test, ans)
	require.NotEmpty(t, test)
	require.NotNil(t, test)
}