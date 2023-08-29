package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagDominicanRepublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 36h31v19H5zm31-19h31v19H36z"/><path fill="#fff" d="M5 33h62v6H5z"/><path fill="#fff" d="M33 17h6v38h-6z"/><circle cx="36" cy="36" r="3" fill="#5c9e31"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}