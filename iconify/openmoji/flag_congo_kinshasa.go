package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCongoKinshasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M5 17h62v38H5z"/><path fill="#d22f27" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M67 24v-7h-8L5 48v7h8l54-31z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="m12.348 30.583l2.323-7l2.003 6.893l-5.564-4.154l7-.172l-5.762 4.433z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}