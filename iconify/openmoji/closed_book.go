package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="m12 58l47.847-.15v-45.7H15.669L12 16z"/><path fill="#d22f27" d="M55.002 61H11.998a.998.998 0 0 1-.998-.998V16.998c0-.551.447-.998.998-.998h43.004c.551 0 .998.447.998.998v43.004a.998.998 0 0 1-.998.998z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m11 17l4.998-5h43.004c.551 0 .998.447.998.998v43.004a.998.998 0 0 1-.998.998"/><path d="M55.002 61H11.998a.998.998 0 0 1-.998-.998V16.998c0-.551.447-.998.998-.998h43.004c.551 0 .998.447.998.998v43.004a.998.998 0 0 1-.998.998z"/></g>`),
		g.Group(children),
	)
}