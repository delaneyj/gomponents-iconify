package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOnSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.654 13.014a6.5 6.5 0 0 0 6.42 5.486h7.852a6.5 6.5 0 1 0 0-13H8.074a6.5 6.5 0 0 0-6.42 7.514ZM16 15.5a3.496 3.496 0 0 0 3.464-3.868A3.496 3.496 0 0 0 16 8.5a3.496 3.496 0 0 0-3.464 3.868A3.496 3.496 0 0 0 16 15.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}