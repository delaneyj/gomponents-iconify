package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoPolymer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M392.3 96h-77.9L160.8 348.6 109.9 256l87.7-160h-77.9L32 256l87.7 160h77.9l153.7-252.6 50.8 92.6-87.7 160h77.9L480 256 392.3 96z" fill="currentColor"/>`),
		g.Group(children),
	)
}