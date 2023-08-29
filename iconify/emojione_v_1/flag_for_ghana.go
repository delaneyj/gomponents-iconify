package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGhana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9cb38" d="M0 25h64v14H0z"/><path fill="#ec1c24" d="M54 10H10C3.373 10 0 14.925 0 21v4h64v-4c0-6.075-3.373-11-10-11"/><path fill="#137a08" d="M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-4H0v4"/><path fill="#25333a" d="m31.778 25.825l1.947 3.945l4.355.638l-3.149 3.072l.744 4.333l-3.897-2.047l-3.896 2.047l.744-4.333l-3.15-3.072l4.356-.638z"/>`),
		g.Group(children),
	)
}