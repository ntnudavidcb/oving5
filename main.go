package main

import(
    "./driver"
    "time"
)

func main(){
    Elev_init()

    for (Elev_get_floor_sensor_signal != 3){
        Elev_set_motor_direction(1)
    }
}

  //  FloorChan := make(chan int,1)
   // ButtonChan :=make(chan Ressurs.MessageFromEvent,4)
    
   // LightChan :=make(chan Ressurs.ButtonLigth,2)
    //MotorChan :=make(chan string,1)
    
   // StatusChan :=make(chan Ressurs.MessageFromEvent,1)
    
    //FromNetwork :=make(chan Ressurs.MessageFromEvent,1)
    //ToNetwork :=make(chan Ressurs.MessageFromEvent,1)
    
   // go elev.DriverInit(ButtonChan, FloorChan,LightChan,MotorChan)
    
    //go elevator.InitElevatorControl(FromNetwork,StatusChan,FloorChan,LightChan,MotorChan,ToNetwork)
    
    //go network.MasterSlaveNetwork(StatusChan,ButtonChan,FromNetwork,ToNetwork)
    
    //for{
    //    time.Sleep(1e9)
    //}