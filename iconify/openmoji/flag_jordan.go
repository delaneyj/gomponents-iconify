package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagJordan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><path d="M5 17h62v13H5z"/><path fill="#d22f27" d="M26 36L5 55V17l21 19z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m13 33.5l.556 1.445l1.449-.455l-.755 1.348l1.25.878l-1.498.234l.111 1.55L13 37.445L11.887 38.5l.111-1.55l-1.498-.234l1.25-.878l-.755-1.348l1.449.455L13 33.5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}