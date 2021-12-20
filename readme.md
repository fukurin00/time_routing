# Time routing provider
This repository is working on synerex platform 

# Usage
- `go build`
- `time_routing`

## Command Line Arguments
- `time_routinig -h`

## Map Format
- You need image file(pgm) and yaml for map and configuration
- the coordination of map is same as ROS (Robot Operating System)
  - watch [here](http://wiki.ros.org/map_server#:~:text=Subscribed%20Topics-,Map%20format,-Maps%20manipulated%20by) 

## Subscribe Messages
### Route Request
- brief: start planning when receiving this message
- [SupplyName](https://github.com/synerex/synerex_api/blob/bf2b537500d3a4d19863340df746eea5fad000c2/synerex.proto#L58): "RouteDemand"
- type: [proto_cav::PathRequest](https://github.com/synerex/proto_cav/blob/f6510ad18d655b5c60653ee2a00bf6cbb8e967b6/cav.proto#L35)
  - int64 robotId: unique id for each agent
  - int64 seq: message sequence number (not used)
  - float radius: radius of agent
  - [Point](https://github.com/synerex/proto_cav/blob/f6510ad18d655b5c60653ee2a00bf6cbb8e967b6/cav.proto#L71) start: start coordination from origin (m)
  - [Point](https://github.com/synerex/proto_cav/blob/f6510ad18d655b5c60653ee2a00bf6cbb8e967b6/cav.proto#L71) goal: goal coordination from origin (m)

## Publish Messages
### Route Supply
- brief: returning path 
- [SupplyName](https://github.com/synerex/synerex_api/blob/bf2b537500d3a4d19863340df746eea5fad000c2/synerex.proto#L58): "SupplyRoute"
- type: [proto_cav::Path](https://github.com/synerex/proto_cav/blob/f6510ad18d655b5c60653ee2a00bf6cbb8e967b6/cav.proto#L46)