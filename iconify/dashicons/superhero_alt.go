package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperheroAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 5H5L2 8l8 8l8-8l-3-3zm-3.3 6.9L10 11l-1.7.9l.3-1.9l-1.4-1.4l1.9-.3l.9-1.7l.9 1.8l1.9.3l-1.4 1.3l.3 1.9z"/>`),
		g.Group(children),
	)
}