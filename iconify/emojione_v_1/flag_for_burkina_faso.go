package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBurkinaFaso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#128807" d="M54 54H10C3.373 54 0 49.075 0 43V32h64v11c0 6.075-3.373 11-10 11"/><path fill="#ec1c24" d="M0 21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v11H0V21z"/><path fill="#f9cb38" d="m41.621 29.07l-7.421.01l-2.3-7.506l-2.285 7.506l-7.425-.01l6.01 4.57l-2.327 7.46l6.04-4.637l6.04 4.637l-2.338-7.46z"/>`),
		g.Group(children),
	)
}