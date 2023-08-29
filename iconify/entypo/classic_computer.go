package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClassicComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 0H4C2.9 0 2 .899 2 2v15a1 1 0 0 0 1 1v2h14v-2a1 1 0 0 0 1-1V2c0-1.101-.899-2-2-2zm-2 15h-4v-1h4v1zm1-4H5V3h10v8z"/>`),
		g.Group(children),
	)
}