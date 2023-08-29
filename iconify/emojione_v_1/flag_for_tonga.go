package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTonga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M54 10H22v22H0v11c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11V21c0-6.075-3.373-11-10-11"/><path fill="#e6e7e8" d="M22 10H10C3.373 10 0 14.925 0 21v11h22V10z"/><path fill="#ec1c24" d="M14 19v-5H8v5H3v6h5v5h6v-5h5v-6z"/>`),
		g.Group(children),
	)
}