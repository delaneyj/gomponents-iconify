package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 65.002L0 711.584h1200L600 65.002zM0 825.623v309.375h1200V825.623H0z"/>`),
		g.Group(children),
	)
}