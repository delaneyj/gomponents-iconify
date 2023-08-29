package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagHonduras(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#61b2e4" d="M5 17h62v13H5zm0 25h62v13H5z"/><circle cx="36" cy="36" r="1.5" fill="#61b2e4"/><circle cx="43" cy="39" r="1.5" fill="#61b2e4"/><circle cx="43" cy="33" r="1.5" fill="#61b2e4"/><circle cx="29" cy="39" r="1.5" fill="#61b2e4"/><circle cx="29" cy="33" r="1.5" fill="#61b2e4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}