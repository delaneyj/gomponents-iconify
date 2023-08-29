package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSwitzerland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M17 17h38v38H17z"/><path fill="#fff" stroke="#fff" stroke-miterlimit="10" stroke-width="2" d="M47 32.462h-7.462V25h-7.076v7.462H25v7.076h7.462V47h7.076v-7.462H47v-7.076z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h38v38H17z"/>`),
		g.Group(children),
	)
}