package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5v14h20V5H2m4 12H4V7h2c2.86.09 5.1 2.33 5 5c.1 2.67-2.14 4.91-5 5m14 0h-2c-2.86-.09-5.1-2.33-5-5c-.1-2.67 2.14-4.91 5-5h2v10Z"/>`),
		g.Group(children),
	)
}