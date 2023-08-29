package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dockright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.083 7.333v16.334h24.833V7.333H3.083zm16.25 13.335H6.083V10.332h13.25v10.336z"/>`),
		g.Group(children),
	)
}