package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTimorLeste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#f1b31c" d="M38 36L5 55V17l33 19z"/><path d="M26 36L5 55V17l21 19z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m15.35 34.05l-4.63 2.44l5.04.67l-3.53-3.53l.67 5.04l2.45-4.62z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}