package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForWallisAndFutuna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M22 11h20v44H22z"/><path fill="#00209f" d="M10 11C3.373 11 0 15.925 0 22v22c0 6.075 3.373 11 10 11h12V11H10z"/><path fill="#ec1c24" d="M54 11H42v44h12c6.627 0 10-4.925 10-11V22c0-6.075-3.373-11-10-11"/>`),
		g.Group(children),
	)
}