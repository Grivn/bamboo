package election

import (
	"fmt"
	"testing"

	"github.com/gitferry/zeitgeber/types"
	"github.com/stretchr/testify/require"
)

func TestRotation_IsLeader(t *testing.T) {
	elect := NewRotation(4)
	leaderID := elect.FindLeaderFor(1)
	require.True(t, elect.IsLeader(leaderID, 1))

	leaderID = elect.FindLeaderFor(4)
	require.Equal(t, "3", leaderID)

	leaderID = elect.FindLeaderFor(3)
	require.True(t, elect.IsLeader(leaderID, 3))
}

func TestRotation_LeaderList(t *testing.T) {
	elect := NewRotation(16)
	byzNo := 5
	var leaderList []int

	for i := 1; i <= 10000; i++ {
		leaderID := elect.FindLeaderFor(types.View(i))
		// if byz, add 0
		if leaderID.Node() <= byzNo {
			leaderList = append(leaderList, 0)
		} else {
			leaderList = append(leaderList, 1)
		}
	}

	fmt.Println(leaderList)
}
