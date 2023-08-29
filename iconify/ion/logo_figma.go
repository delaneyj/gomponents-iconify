package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoFigma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 176a80 80 0 0 0 0-160H176a80 80 0 0 0 0 160a80 80 0 0 0 0 160a80 80 0 1 0 80 80V176Z"/><circle cx="336" cy="256" r="80" fill="currentColor"/>`),
		g.Group(children),
	)
}