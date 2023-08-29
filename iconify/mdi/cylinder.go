package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cylinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c-2.11 0-9 .29-9 4v12c0 3.71 6.89 4 9 4c2.11 0 9-.29 9-4V6c0-3.71-6.89-4-9-4m0 8c1.38 0 4.8-.13 7-1.25v6.5C16.8 14.13 13.38 14 12 14s-4.8.13-7 1.25v-6.5C7.2 9.87 10.62 10 12 10m0-6c3.87 0 7 .89 7 2c0 1.11-3.13 2-7 2s-7-.89-7-2c0-1.11 3.13-2 7-2m0 16c-3.87 0-7-.89-7-2c0-1.1 3.13-2 7-2s7 .9 7 2c0 1.11-3.13 2-7 2Z"/>`),
		g.Group(children),
	)
}