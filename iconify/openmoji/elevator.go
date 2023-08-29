package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M22 19h28v34H22z"/><path fill="#FFF" d="m29 53l19-34H22v34z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M22 19h28v34H22zm-6 42V11m20 38V23m20 38V11"/>`),
		g.Group(children),
	)
}