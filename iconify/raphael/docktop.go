package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.916 23.667V7.333H3.083v16.334h24.833zm-3-3H6.082v-6.5h18.833v6.5z"/>`),
		g.Group(children),
	)
}