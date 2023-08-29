package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMadagascar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#156d08" d="M4 43c0 6.075 1.373 11 8 11h42c6.627 0 10-4.925 10-11V32H2l2 11z"/><path fill="#c32129" d="M54 10H12C5.373 10 1 14.925 1 21v12h63V21c0-6.075-3.373-11-10-11"/><path fill="#e6e7e8" d="M10 10C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h12V10H10z"/>`),
		g.Group(children),
	)
}