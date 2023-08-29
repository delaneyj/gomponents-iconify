package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBelarus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 42h62v13H5z"/><path fill="#fff" d="M5 17h12v38H5z"/><path fill="none" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 22l3.5-3.5L15 22l-3.5 3.5zm0 7l3.5-3.5L15 29l-3.5 3.5zm0 7l3.5-3.5L15 36l-3.5 3.5zm0 7l3.5-3.5L15 43l-3.5 3.5zm0 7l3.5-3.5L15 50l-3.5 3.5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}