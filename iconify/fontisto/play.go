package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.001 1.165v21.669a1.275 1.275 0 0 0 1.891 1.017l-.006.003l21.442-10.8a1.172 1.172 0 0 0 .007-2.113l-.007-.003L1.886.138A1.273 1.273 0 0 0 .003 1.162v.004z"/>`),
		g.Group(children),
	)
}