# Domain Driven Design in Go

Base objects for DDD:

| Type        | Notes                                                                                   |
| ----------- | --------------------------------------------------------------------------------------- |
| `Version`   | Optimitic lock                                                                          |
| `Entity`    | Base **Entity** with `ID` field                                                         |
| `Aggregate` | Base **Aggregate** extending `Entity` with `Version` field                              |
| `Stream`    | Base **Event Stream Aggregate** extending `Aggregate` with `Events` and `Outbox` fields |
