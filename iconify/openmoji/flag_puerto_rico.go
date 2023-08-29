package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPuertoRico(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 32h62v8H5zm0 15h62v8H5zm0-30h62v8H5z"/><path fill="#1e50a0" d="M26 36L5 55V17l21 19z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m13.5 33.5l1.545 5L11 35.41h5l-4.045 3.09l1.545-5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}