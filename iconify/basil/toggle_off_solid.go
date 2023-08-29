package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22.346 13.014a6.5 6.5 0 0 1-6.42 5.486H8.074a6.5 6.5 0 1 1 0-13h7.852a6.5 6.5 0 0 1 6.42 7.514ZM8 15.5a3.496 3.496 0 0 1-3.464-3.868A3.496 3.496 0 0 1 8 8.5a3.496 3.496 0 0 1 3.464 3.868A3.496 3.496 0 0 1 8 15.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}