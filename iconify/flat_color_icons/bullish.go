package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#4CAF50" d="M40 21h4v23h-4zm-6 7h4v16h-4zm-6-5h4v21h-4zm-6 6h4v15h-4zm-6 3h4v12h-4zm-6-2h4v14h-4zm-6 4h4v10H4z"/><g fill="#388E3C"><path d="M40.1 9.1L34 15.2l-4-4l-10 10l-5-5L4.6 26.6l2.8 2.8l7.6-7.6l5 5l10-10l4 4l8.9-8.9z"/><path d="M44 8h-9l9 9z"/></g>`),
		g.Group(children),
	)
}