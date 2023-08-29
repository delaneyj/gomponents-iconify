package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForNorway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M64 43c0 6.075-3.373 11-10 11H10C3.373 54 0 49.075 0 43V21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v22"/><path fill="#e6e7e8" d="M18 10h11v44H18z"/><path fill="#e6e7e8" d="M0 27h64v10H0z"/><path fill="#2b3990" d="M20 10h7v44h-7z"/><path fill="#2b3990" d="M0 29h64v6H0z"/>`),
		g.Group(children),
	)
}