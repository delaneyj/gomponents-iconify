package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19V8H5v11h14M16 1h2v2h1a2 2 0 0 1 2 2v14c0 1.11-.89 2-2 2H5a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h1V1h2v2h8V1m-4 16.17l-.58-.53C9.36 14.77 8 13.54 8 12.03c0-1.23.97-2.2 2.2-2.2c.7 0 1.36.32 1.8.83c.44-.51 1.1-.83 1.8-.83c1.23 0 2.2.97 2.2 2.2c0 1.51-1.36 2.74-3.42 4.61l-.58.53Z"/>`),
		g.Group(children),
	)
}