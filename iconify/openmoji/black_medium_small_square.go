package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackMediumSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M51.736 20.306H20.264v31.471h31.472V20.306Z"/><path fill="#3F3F3F" d="M51.736 20.306H20.264v31.471h31.472V20.306Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M51.736 20.306H20.264v31.471h31.472V20.306Z"/>`),
		g.Group(children),
	)
}