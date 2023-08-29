package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2zm2.502 5.5H5.5a.5.5 0 0 0 0 1h5.002a.5.5 0 1 0 0-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}