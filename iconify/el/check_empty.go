package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v1200h1200V0H0zm196.875 196.875h806.25v806.25h-806.25v-806.25z"/>`),
		g.Group(children),
	)
}