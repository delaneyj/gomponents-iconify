package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpThinCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.97c-4.41 0-8.03 3.62-8.03 8.03c0 4.41 3.62 8.03 8.03 8.03c4.41 0 8.03-3.62 8.03-8.03c0-4.41-3.62-8.03-8.03-8.03M12 2c5.54 0 10 4.46 10 10s-4.46 10-10 10S2 17.54 2 12S6.46 2 12 2m1 8.46h3L12 6.5l-4 3.96h3v7.04h2"/>`),
		g.Group(children),
	)
}