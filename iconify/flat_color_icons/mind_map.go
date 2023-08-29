package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MindMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#CFD8DC" d="m39.4 23l-.8-4L26 21.6V8h-4v12.3l-13.9-9l-2.2 3.4l15.2 9.8L9.4 39.8l3.2 2.4l11.3-14.8l8.4 12.7l3.4-2.2l-8.4-12.5z"/><circle cx="24" cy="24" r="7" fill="#3F51B5"/><g fill="#00BCD4"><circle cx="24" cy="8" r="5"/><circle cx="39" cy="21" r="5"/><circle cx="7" cy="13" r="5"/><circle cx="11" cy="41" r="5"/><circle cx="34" cy="39" r="5"/></g>`),
		g.Group(children),
	)
}