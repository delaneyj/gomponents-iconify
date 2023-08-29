package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NegativeDynamic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M19 22h10v20H19zM6 8h10v34H6zm26 22h10v12H32z"/><g fill="#3F51B5"><path d="M42 12L32 22h10z"/><path d="m27.561 10.396l2.828-2.828l9.969 9.969l-2.828 2.828z"/></g>`),
		g.Group(children),
	)
}