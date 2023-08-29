package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dockbottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.083 7.333v16.334h24.833V7.333H3.083zm21.832 9.5H6.083v-6.5h18.833v6.5z"/>`),
		g.Group(children),
	)
}