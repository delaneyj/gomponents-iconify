package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMauritius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M0 21h64c0-6.075-3.373-11-10-11H10C3.373 10 0 14.925 0 21z"/><path fill="#2b3990" d="M0 21h64v11H0z"/><path fill="#f9cb38" d="M0 32h64v11H0z"/><path fill="#137a08" d="M64 43H0c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11"/>`),
		g.Group(children),
	)
}