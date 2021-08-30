package model

import "strconv"

func (rl RoomLocation) toString() string {
	name := ""

	//parse direction first

	name += i2D(rl.room_corridor)

	name += i2F(rl.room_floor_level)

	name += i2Rn(rl.room_number)

	return name
}

//****************************************************************************************************************

//Helpers

func i2F(i int) string {
	return strconv.Itoa(i)
}

func i2Rn(i int) string {
	return strconv.Itoa(i)
}

func i2D(i int) string {
	if i == __ORIENATION_NORTH {
		return "Z" + __ORIENATION_NORTH_S
	}
	if i == __ORIENATION_EAST {
		return "Z" + __ORIENATION_EAST_S
	}
	if i == __ORIENATION_SOUTH {
		return "Z" + __ORIENATION_SOUTH_S
	}
	if i == __ORIENATION_WEST {
		return "Z" + __ORIENATION_WEST_S
	}
	return ""
}

func D2i(e string) int {
	if e == __ORIENATION_WEST_S {
		return __ORIENATION_WEST
	}
	if e == __ORIENATION_NORTH_S {
		return __ORIENATION_NORTH
	}
	if e == __ORIENATION_EAST_S {
		return __ORIENATION_EAST
	}
	if e == __ORIENATION_SOUTH_S {
		return __ORIENATION_SOUTH
	}
	return 0
}
