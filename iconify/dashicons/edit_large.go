package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.4 14.1l1.3 1.3l6.9-6.9l-1.3-1.3l-6.9 6.9zm6.3-7.5l-1.3-1.3l-6.9 6.9l1.4 1.4l6.8-7zm2.1-4.7l3.3 3.3c.6.6.5 1.5 0 2l-9.9 9.9l-6.9 1.4l1.4-6.9c6.2-6.3 9.5-9.6 9.9-9.9c.6-.4 1.6-.4 2.2.2z"/>`),
		g.Group(children),
	)
}