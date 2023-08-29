package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCrown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9 4v6c0 5.55-3.84 10.74-9 12c-5.16-1.26-9-6.45-9-12V5l9-4m4 13H8v1.5c0 .27.19.46.47.5h6.96c.31 0 .52-.16.57-.41V14m1-6l-2.67 2.67L12 8.34l-2.33 2.33L7 8l1 5h8l1-5Z"/>`),
		g.Group(children),
	)
}