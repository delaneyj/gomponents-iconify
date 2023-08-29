package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMalawi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path d="M5 17h62v13H5z"/><path fill="#d22f27" d="M5 30h62v12H5z"/><path fill="#d22f27" stroke="#d22f27" stroke-miterlimit="10" d="M29.24 31a7.503 7.503 0 0 1 13.52-.002Z"/><path fill="none" stroke="#d22f27" stroke-miterlimit="10" d="M26.54 31a10.004 10.004 0 0 1 18.92 0Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}