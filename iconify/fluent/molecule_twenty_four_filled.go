package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoleculeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 12a5 5 0 1 0-4.337-2.51l-2.714 1.808a4 4 0 1 0 .23 5.13l3.887 1.943a3 3 0 1 0 .672-1.341L9.85 15.087a4.004 4.004 0 0 0-.113-2.513l2.863-1.907A4.982 4.982 0 0 0 16 12Z"/>`),
		g.Group(children),
	)
}