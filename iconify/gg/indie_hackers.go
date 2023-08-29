package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndieHackers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6h3v12H4V6Zm6 0h3v4.5h4V6h3v12h-3v-4.5h-4V18h-3V6Z"/>`),
		g.Group(children),
	)
}