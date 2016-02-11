package driver
/*
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
*/
//import "C"

import (
	"io.go"
	"channels.go"
)

const N_FLOORS = 4
const N_BUTTONS = 3

type ELEV_BUTTON_TYPE int
const (
    CALL_UP = 0
    CALL_DOWN = 1
    CALL_COMMAND = 2
)

type ELEV_MOTOR_DIR int
const (
    DIR_UP = 1
    DIR_DOWN = -1
    DIR_STOP = 0
)

var lamp_channel_matrix =[N_FLOORS][N_BUTTONS]int{
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var button_channel_matrix =[N_FLOORS][N_BUTTONS]int{
    {BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    {BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    {BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    {BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}


func Elev_init() bool{
	if io_init()==0{
        return false
    }
    for f=0;f<N_FLOORS;i++{
    	for (ELEV_BUTTON_TYPE b = 0; b < N_BUTTONS; b++){
    		Elev_set_button_lamp(b,f,0)
    	}
    }
}

func Elev_set_motor_direction(dirn ELEV_MOTOR_DIR){ 
	if (dirn == 0){
        Io_write_analog(MOTOR, 0)
    } else if (dirn > 0) {
        Io_clear_bit(MOTORDIR)
        Io_write_analog(MOTOR, MOTOR_SPEED)
    } else if (dirn < 0) {
        Io_set_bit(MOTORDIR)
        Io_write_analog(MOTOR, MOTOR_SPEED)
    }
}

func Elev_set_button_lamp(button ELEV_BUTTON_TYPE, floor int, value int){
	if (value) {
        Io_set_bit(lamp_channel_matrix[floor][button]);
    } else {
        Io_clear_bit(lamp_channel_matrix[floor][button]);
    }
}

func Elev_set_floor_indicator(floor int){
	// Binary encoding. One light must always be on.
    if (floor & 0x02) {
        Io_set_bit(LIGHT_FLOOR_IND1); //finn ut hvor LIGHT_FLOOR_IND1 ligger
    } else {
        Io_clear_bit(LIGHT_FLOOR_IND1);
    }    

    if (floor & 0x01) {
        Io_set_bit(LIGHT_FLOOR_IND2);
    } else {
        Io_clear_bit(LIGHT_FLOOR_IND2);
    }   
}

func Elev_set_door_open_lamp(value int){
	if (value) {
        Io_set_bit(LIGHT_DOOR_OPEN); //samme som i funksjonen over
    } else {
        Io_clear_bit(LIGHT_DOOR_OPEN);
    }
}

func Elev_set_stop_lamp(value int){
	if (value) {
        Io_set_bit(LIGHT_STOP); //samme som i funksjonen over
    } else {
        Io_clear_bit(LIGHT_STOP);
    }
}

func Elev_get_button_signal(button BUTTON_TYPE, floor int) int {
	return int(C.elev_get_button_signal(button, C.int(floor)))
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