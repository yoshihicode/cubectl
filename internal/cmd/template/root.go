package template

var RootHelpTemplate = `cubectl controls a cube (and a pea pod) instead of Kubernetes clusters.

Find more information at:
  https://github.com/yoshihicode/cubectl

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Basic Cube Commands (Beginner):
  delete      Delete a resource
  get         Display a resource

Troubleshooting and Debugging Commands:
  describe    Describe a resource
  logs        Print the logs for a resource

Options:
  -h, --help            help for cubectl
  -o, --output string   Output format: wireframe|solid (default "wireframe")
  -w, --watch           Watch for changes to the cube (it will keep spinning)

Usage:
  cubectl [command] [options]

Use "cubectl <command> --help" for more information about a command.
`

var RootUsageTemplate = `cubectl controls a cube (and a pea pod) instead of Kubernetes clusters.

Find more information at:
  https://github.com/yoshihicode/cubectl

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Basic Cube Commands (Beginner):
  delete      Delete a resource
  get         Display a resource

Troubleshooting and Debugging Commands:
  describe    Describe a resource
  logs        Print the logs for a resource

Options:
  -h, --help            help for cubectl
  -o, --output string   Output format: wireframe|solid (default "wireframe")
  -w, --watch           Watch for changes to the cube (it will keep spinning)

Usage:
  cubectl [command] [options]

Use "cubectl <command> --help" for more information about a command.
`
