package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagNamibia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17v38l62-38H5z"/><path fill="#d22f27" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M67 24v-7h-8L5 48v7h8l54-31z"/><circle cx="17" cy="27" r="4" fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}