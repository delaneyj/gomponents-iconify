package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 15.5a3.496 3.496 0 0 1-3.464-3.868A3.496 3.496 0 0 1 8 8.5a3.496 3.496 0 0 1 3.464 3.868A3.496 3.496 0 0 1 8 15.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M8.074 18.75h7.852a6.75 6.75 0 0 0 0-13.5H8.074a6.75 6.75 0 0 0 0 13.5Zm0-1.5a5.25 5.25 0 1 1 0-10.5h7.852a5.25 5.25 0 1 1 0 10.5H8.074Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}