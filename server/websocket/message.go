package websocket

import (
	"encoding/json"
	"log"
)
const(
	 CmdType_HEARTBEAT = iota+1
	 CmdType_KILL

)


type RobotMsg struct {
	Cmd   int      `json:"cmd"`
	Data  string `json:"data"`
}


func NewRobotMsg(cmd int,data string) *RobotMsg{
	return &RobotMsg{cmd,data}
}

func (r *RobotMsg)ToBytes() ([]byte,error){
	msg, err := json.Marshal(r)
	if err != nil {
		log.Printf("Fail to package robotMsg :%v", err)
		return nil,err
	}
	return msg,nil
}



func ParseRobotMsg(message []byte) *RobotMsg {
	var data RobotMsg
	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Println("Fail to parse message:%v", err)
		return nil
	}
	return &data
}

func PackageRobotMsg(cmd int,   data string) []byte {
	var req = RobotMsg{
		Cmd:   cmd,
		Data:  data,
	}

	msg, err := json.Marshal(req)
	if err != nil {
		log.Println("Fail to packageResponseMsg:%v", err)
		return nil
	}
	return msg
}
