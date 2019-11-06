# ecs-overrides

WIP

This is a CLI tool for printing ecs overrides JSON string.
It is supposed to be used on executing [ecs run-task command](https://docs.aws.amazon.com/cli/latest/reference/ecs/run-task.html).

## Usage

### cmd [name] [command]

Only override container name and command.

```sh
$ ecs-overrides cmd app "echo hello world"
{"ContainerOverrides":[{"Command":["echo","hello","world"],"Cpu":null,"Environment":null,"Memory":null,"MemoryReservation":null,"Name":"app","ResourceRequirements":null}],"ExecutionRoleArn":null,"InferenceAcceleratorOverrides":null,"TaskRoleArn":null}%
```

