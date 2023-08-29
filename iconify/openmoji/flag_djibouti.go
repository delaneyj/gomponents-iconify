package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagDjibouti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 36h62v19H5z"/><path fill="#fff" d="M26 36L5 55V17l21 19z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m11.413 39.046l2.323-7l2.003 6.893l-5.564-4.154l7-.172l-5.762 4.433z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}