package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PositiveDynamic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M19 22h10v20H19zM32 8h10v34H32zM6 30h10v12H6z"/><g fill="#3F51B5"><path d="m11 8l10 10V8z"/><path d="m9.394 22.437l-2.828-2.828l9.969-9.969l2.828 2.828z"/></g>`),
		g.Group(children),
	)
}