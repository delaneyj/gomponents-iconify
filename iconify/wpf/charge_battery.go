package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargeBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M23 5H5C3.3 5 2 6.3 2 8v1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h1v1c0 1.7 1.3 3 3 3h18c1.7 0 3-1.3 3-3V8c0-1.7-1.3-3-3-3zm-1 9l-7-1v4l-8-5l7 1V9l8 5z"/>`),
		g.Group(children),
	)
}