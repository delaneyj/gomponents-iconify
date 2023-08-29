package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForUkraine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M54 10H10C3.373 10 0 14.925 0 21v11h64V21c0-6.075-3.373-11-10-11"/><path fill="#f9cb38" d="M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11V32H0v11z"/>`),
		g.Group(children),
	)
}