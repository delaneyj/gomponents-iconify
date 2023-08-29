package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM1 4a3 3 0 0 1 5.83-1h10.34A3.001 3.001 0 1 1 21 6.83v10.34A3.001 3.001 0 1 1 17.17 21H6.83A3.001 3.001 0 1 1 3 17.17V6.83A3.001 3.001 0 0 1 1 4Zm4 2.83v10.34A3.008 3.008 0 0 1 6.83 19h10.34A3.009 3.009 0 0 1 19 17.17V6.83A3.008 3.008 0 0 1 17.17 5H6.83A3.008 3.008 0 0 1 5 6.83ZM20 3a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM4 19a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm16 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}