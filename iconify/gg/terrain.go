package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terrain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 10l-5 8h10l-5-8Zm2.529.754L13.5 6L21 18h-5.943l-4.528-7.246Z"/>`),
		g.Group(children),
	)
}