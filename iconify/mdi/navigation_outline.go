package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7.3l4.3 10.4l-3.5-1.5l-.8-.4l-.8.4l-3.5 1.5L12 7.3M12 2L4.5 20.3l.7.7l6.8-3l6.8 3l.7-.7L12 2Z"/>`),
		g.Group(children),
	)
}