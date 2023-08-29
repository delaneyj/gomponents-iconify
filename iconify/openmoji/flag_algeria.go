package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAlgeria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#fff" d="M36 17h31v38H36z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m43.863 38.5l-2.406-4.572l4.577 2.398l-5 .692l3.513-3.518l-.684 5z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M39.292 44.643a8.643 8.643 0 1 1 3.958-16.335a11 11 0 1 0 0 15.384a8.715 8.715 0 0 1-3.958.95Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}