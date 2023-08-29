package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M22 19h28v34H22z"/><path fill="#FFF" d="M29 53l19-34H22v34z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M22 19h28v34H22z"/><path d="M16 61V11"/><path d="M36 49V23"/><path d="M56 61V11"/></g>`),
		g.Group(children),
	)
}