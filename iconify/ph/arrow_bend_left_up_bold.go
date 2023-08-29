package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 224a12 12 0 0 1-12 12A108.12 108.12 0 0 1 92 128V61L64.49 88.49a12 12 0 0 1-17-17l48-48a12 12 0 0 1 17 0l48 48a12 12 0 0 1-17 17L116 61v67a84.09 84.09 0 0 0 84 84a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}