package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFrenchGuiana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M67 17v38L5 17h62z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m33.703 39.269l2.539-7.446l2.189 7.332l-6.081-4.418l7.65-.184l-6.297 4.716z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}