package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M80 352h64v64H80z" fill="currentColor"/><path d="M176 288h64v128h-64z" fill="currentColor"/><path d="M272 192h64v224h-64z" fill="currentColor"/><path d="M368 96h64v320h-64z" fill="currentColor"/>`),
		g.Group(children),
	)
}