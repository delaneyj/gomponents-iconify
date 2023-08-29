package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForIndonesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M54 54H10C3.373 54 0 49.075 0 43V32h64v11c0 6.075-3.373 11-10 11"/><path fill="#ec1c24" d="M0 21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v11H0V21z"/>`),
		g.Group(children),
	)
}