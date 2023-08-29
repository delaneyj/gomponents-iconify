package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagZambia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#f1b31c" d="M61 36h6v19h-6z"/><path fill="#d22f27" d="M51 36h5v19h-5z"/><path d="M56 36h5v19h-5z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m53 31l5.75 1L64 31"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M58 32v2m1-2v2m-.5-3.5v2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}