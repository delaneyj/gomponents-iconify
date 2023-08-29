package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-3.25 16h-1.5V8.9L6.72 18H4.66L15.45 6.5c.21-.24.55-.31.82-.2c.29.12.48.39.48.7v11Z"/>`),
		g.Group(children),
	)
}