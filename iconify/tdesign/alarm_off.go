package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m.086 6.5l2.5-2.5l-2-2L2 .586l2 2l1.5-1.5L6.914 2.5L5.414 4l18 18L22 23.414l-3.139-3.139A9.968 9.968 0 0 1 12.001 23C6.476 23 2 18.523 2 13a9.968 9.968 0 0 1 2.724-6.861L4 5.414l-2.5 2.5L.086 6.5ZM6.14 7.554A8 8 0 0 0 17.446 18.86L6.14 7.554Zm2.421-4l.97-.246A9.995 9.995 0 0 1 12 3c5.523 0 10 4.477 10 10c0 .851-.107 1.679-.308 2.47l-.246.969l-1.938-.493l.246-.969a8 8 0 0 0-9.731-9.731l-.97.246l-.492-1.938Zm9.94-2.468L23.913 6.5L22.5 7.914L17.086 2.5L18.5 1.086Z"/>`),
		g.Group(children),
	)
}