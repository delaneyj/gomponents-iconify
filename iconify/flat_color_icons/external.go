package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func External(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="31" r="14" fill="#B2DFDB"/><g fill="#009688"><path d="M24 3.3L33 14H15z"/><path d="M21 11h6v23h-6z"/></g>`),
		g.Group(children),
	)
}