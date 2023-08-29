package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonTriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.027 8a8 8 0 0 0-8-8a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8zM5.154 10c-.205-.186-.205-.774 0-.96l3.467-4.9a.566.566 0 0 1 .746 0l3.478 4.9c.207.185.207.773 0 .96h-7.69z"/>`),
		g.Group(children),
	)
}