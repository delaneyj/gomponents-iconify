package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComboChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M37 18h6v24h-6zm-8 8h6v16h-6zm-8-4h6v20h-6zm-8 10h6v10h-6zm-8-4h6v14H5z"/><g fill="#3F51B5"><circle cx="8" cy="16" r="3"/><circle cx="16" cy="18" r="3"/><circle cx="24" cy="11" r="3"/><circle cx="32" cy="13" r="3"/><circle cx="40" cy="9" r="3"/><path d="m39.1 7.2l-7.3 3.7l-8.3-2.1l-8 7l-7-1.7l-1 3.8l9 2.3l8-7l7.7 1.9l8.7-4.3z"/></g>`),
		g.Group(children),
	)
}