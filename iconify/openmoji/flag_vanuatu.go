package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagVanuatu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 36h62v19H5z"/><path d="M26 36L5 55V17l21 19z"/><path fill="#f1b31c" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M26 34L5 17v4l18.529 15L5 51v4l21-17h41v-4H26z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M12.333 39.273a3.929 3.929 0 0 0 4.167-3.637a3.143 3.143 0 0 0-3.333-2.909a2.514 2.514 0 0 0-2.667 2.328a2.012 2.012 0 0 0 2.133 1.861"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}