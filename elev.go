package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.h and driver.go
/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
*/

import "C"

func Elev_init(){
	C.elev_init()
}

func Elev_set_motor_direction(dirn C.struct_elev_motor_direction_t){ //dobbeltsjekk om C.struct... er en god løsning for å importere typedef's fra C
	C.elev_set_motor_direction()
}

func Elev_set_button_lamp(button C.struct_elev_button_type_t, floor int, value int){
	C.elev_set_button_lamp(button elev_button_type_t, C.int(floor), C.int(value))
}

func Elev_set_floor_indicator(floor int){
	C.elev_set_floor_indicator(C.int(floor))
}

func Elev_set_door_open_lamp(int value){
	C.elev_set_door_open_lamp(C.int(value))
}

func Elev_set_stop_lamp(value int){
	C.elev_set_stop_lamp(C.int(value))
}

func Elev_get_button_signal(button C.struct_elev_button_type_t, floor int) int {
	return int(C.elev_get_button_signal(button elev_button_type_t, C.int(floor)))
}

func Elev_get_floor_sensor_signal() int {
	return int(C.elev_get_floor_sensor_signal())
}

func Elev_get_stop_signal() int{
	return int(C.elev_get_stop_signal())
}

func Elev_get_obstruction_signal() int{
	return int(C.elev_get_obstruction_signal())
}