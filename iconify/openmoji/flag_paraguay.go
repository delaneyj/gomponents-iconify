package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagParaguay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v13H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="m34.707 38l1.328-4l1.145 3.939L34 35.565l4-.098L34.707 38z"/><circle cx="36" cy="36" r="5" fill="none" stroke="#5c9e31" stroke-miterlimit="10"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}