package menu

import (
	"mus_projekt/app/controller/rooms"
	"mus_projekt/app/model"
	"strconv"
)

const BUILDING_FLOOR_COUNT string = "number_of_floors"
const BUILDING_FLOOR_HEAD string = "f"

const BUILDING_FLOOR_NAME string = "floor_name"
const BUILDING_FLOOR string = "floor"

var BUILDING_CORRIDORS []string = []string{"N", "E", "S", "W"}

const BUILDING_CORRIDOR_DIRECTION string = "direction"

var BUILDING_CORRIDOR_DIRECTIONS []string = []string{"North", "East", "South", "West"}

const BUILDING_CORRIDOR_NAME string = "corridor_name"

const BULDING_CORRIDOR_ROOMS string = "rooms"

const ROOM_NAME string = "room_name"
const ROOM_NAME_MISC string = "room_name_misc"
const ROOM_ID string = "room_id"

const FLOOR_NAME string = "Level"

func CreateAPIData() map[string]interface{} {

	m := make(map[string]interface{})

	m[BUILDING_FLOOR_COUNT] = rooms.BUILDING_FLOORS
	var i int
	for i = 0; i < rooms.BUILDING_FLOORS; i++ {
		m[BUILDING_FLOOR_HEAD+strconv.Itoa(i)] = createFloor(i)
	}
	return m
}

func createFloor(i int) map[string]interface{} {
	m := make(map[string]interface{})

	m[BUILDING_FLOOR_NAME] = ""
	m[BUILDING_FLOOR] = FLOOR_NAME + " " + strconv.Itoa(i)
	for j, k := range BUILDING_CORRIDORS {

		m[k] = createCorridor(i, j) //Necessary OFFSET N = 1, E = 2, S = 3; W = 4
	}
	return m
}

func createCorridor(i int, j int) map[string]interface{} {
	m := make(map[string]interface{})

	m[BUILDING_CORRIDOR_DIRECTION] = BUILDING_CORRIDOR_DIRECTIONS[j]
	m[BUILDING_CORRIDOR_NAME] = ""
	m[BULDING_CORRIDOR_ROOMS] = createCorridorRooms(i, j)
	return m
}

func createCorridorRooms(i int, j int) []map[string]interface{} {
	rms := model.GetAllRoomsByInt(i, j+1)
	mps := []map[string]interface{}{}

	var rnbr = 0
	for rnbr = 0; rnbr < rooms.BUILDING_FLOORS; rnbr++ {
		for _, r := range *rms {
			if r.RoomNumber() == rnbr {
				m := make(map[string]interface{})
				m[ROOM_NAME] = r.Location().Name()
				m[ROOM_NAME_MISC] = r.Room_Name_Misc()
				m[ROOM_ID] = r.ID()
				mps = append(mps, m)
				break
			}
		}
	}
	//mps = append(mps,make(map[string]interface{}))
	return mps
}
