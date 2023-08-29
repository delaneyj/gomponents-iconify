package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSouthAfrica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 36h62v19H5z"/><path fill="#5c9e31" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 55l21-13h41V30H26L5 17v38z"/><path stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 36L5 46V26l15 10z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}