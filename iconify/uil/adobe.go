package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.07 17.28h2.78l1.75 3.44h2.54L12 9.87ZM2 3v18L9.42 3Zm12.48 0L22 20.81V3Z"/>`),
		g.Group(children),
	)
}