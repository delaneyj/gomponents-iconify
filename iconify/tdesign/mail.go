package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v1.83l9 4.55l9-4.55V5H3Zm18 4.07l-9 4.55l-9-4.55V19h18V9.07Z"/>`),
		g.Group(children),
	)
}