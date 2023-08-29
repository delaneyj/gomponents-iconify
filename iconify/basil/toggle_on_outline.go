package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 15.5a3.496 3.496 0 0 0 3.464-3.868A3.496 3.496 0 0 0 16 8.5a3.496 3.496 0 0 0-3.464 3.868A3.496 3.496 0 0 0 16 15.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M15.926 18.75H8.074a6.75 6.75 0 0 1 0-13.5h7.852a6.75 6.75 0 0 1 0 13.5Zm0-1.5a5.25 5.25 0 1 0 0-10.5H8.074a5.25 5.25 0 0 0 0 10.5h7.852Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}