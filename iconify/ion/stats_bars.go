package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatsBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M176 96h64v320h-64z" fill="currentColor"/><path d="M80 320h64v96H80z" fill="currentColor"/><path d="M272 256h64v160h-64z" fill="currentColor"/><path d="M368 192h64v224h-64z" fill="currentColor"/>`),
		g.Group(children),
	)
}