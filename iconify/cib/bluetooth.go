package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20.288 10.694l-2.681 2.681L17.588 8l2.7 2.694zm-2.701 13.737l2.694-2.694l-2.681-2.681zM28 16.212C28 29.062 23.506 32 16.431 32S4 29.062 4 16.212S9.212 0 16.288 0S28 3.369 28 16.212zm-9.906 0l4.962-5.537l-7.819-8.394v11.056l-4.613-4.612l-1.688 1.681l5.794 5.812l-5.794 5.813l1.681 1.681L15.23 19.1l.144 10.625l7.962-7.969z"/>`),
		g.Group(children),
	)
}