# ecs-overrides

WIP

This is a CLI tool for printing ecs overrides JSON string.
It is supposed to be used on executing [ecs run-task command](https://docs.aws.amazon.com/cli/latest/reference/ecs/run-task.html).

## Usage

### [--env <KEY=VAL>] [name] [command]

override container name and command, and environment.

* Basic

```sh
$ ecs-overrides app echo hello world
{"ContainerOverrides":[{"Command":["echo","hello","world"],"Cpu":null,"Environment":null,"Memory":null,"MemoryReservation":null,"Name":"app","ResourceRequirements":null}],"ExecutionRoleArn":null,"InferenceAcceleratorOverrides":null,"TaskRoleArn":null}%
```

* With environment

```sh
$ ecs-overrides -e XXX=abc -e YYY=def app echo hello world
{"ContainerOverrides":[{"Command":["echo","hello","world"],"Cpu":null,"Environment":[{"Name":"XXX","Value":"abc"},{"Name":"YYY","Value":"def"}],"Memory":null,"MemoryReservation":null,"Name":"app","ResourceRequirements":null}],"ExecutionRoleArn":null,"InferenceAcceleratorOverrides":null,"TaskRoleArn":null}%
```
