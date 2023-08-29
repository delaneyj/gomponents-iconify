package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.883 21.095l9.098-3.02L22 3H4.958L4.31 6.364h13.657l-.425 2.148H3.864l-.669 3.363h13.658l-.75 3.85l-5.511 1.803l-4.762-1.802l.325-1.682H2.79L2 18.076l7.883 3.02"/>`),
		g.Group(children),
	)
}