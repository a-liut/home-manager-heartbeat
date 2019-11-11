# home-manager-heartbeat
Tool that checks health status of home-manager registered devices.

## How does it works

The list of registered devices is retrieved from the *home-manager-devices* service each 20 seconds:

```GET http://<devicesUrl>/devices```

Then, each device is contacted to their respective HeartbeatUrl. The device is online if the request can be completed successfully.

The device description is updated only if the online status changes.
