package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.13 104.13 0 0 0 128 24ZM84 140a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm44 0a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm44 0a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}