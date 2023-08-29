package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeteorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 1v12A9 9 0 1 1 7.375 5.278L14 1.453v2.77L21 1Zm-9 7a5 5 0 1 0 0 10a5 5 0 0 0 0-10Z"/>`),
		g.Group(children),
	)
}