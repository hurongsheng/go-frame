//有 N 个房间，开始时你位于 0 号房间。每个房间有不同的号码：0，1，2，...，N-1，并且房间里可能有一些钥匙能使你进入下一个房间。 
package slicerange

import "fmt"

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]struct{}, 0)
	visited = visitRooms(rooms, visited, 0)
	fmt.Println(visited)
	return len(visited) == len(rooms)
}

func visitRooms(rooms [][]int, visited map[int]struct{}, roomId int) map[int]struct{} {
	if _, ok := visited[roomId]; ok {
		return visited
	}
	visited[roomId] = struct{}{}
	for _, key := range rooms[roomId] {
		if _, ok := visited[key]; !ok {
			visited = visitRooms(rooms, visited, key)
			fmt.Println(visited)
		}
	}
	return visited
}

//leetcode submit region end(Prohibit modification and deletion)
