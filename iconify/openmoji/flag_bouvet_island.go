package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBouvetIsland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-width="2"><path fill="#d22f27" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M5 17h62v38H5z"/><path fill="#1e50a0" stroke="#fff" stroke-miterlimit="10" d="M67 33H30V17h-6v16H5v6h19v16h6V39h37v-6z"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}