package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupWithStraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D22F27" d="m49.923 50.25l1.977-14.4H20.4l2.259 14.4z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path d="m49.923 50.25l1.977-14.4H20.4l2.259 14.4z"/><path stroke-linecap="round" stroke-linejoin="round" d="m54 20.55l-6.3 45.9H25.2L18 20.55m0 0h36M36 5v12"/></g>`),
		g.Group(children),
	)
}