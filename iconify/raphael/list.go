package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.082 4.083v3h22.835v-3H4.082zm0 16.223h22.835v-3H4.082v3zm0-6.612h22.835v-3H4.082v3zm0 13.223h22.835v-3H4.082v3z"/>`),
		g.Group(children),
	)
}