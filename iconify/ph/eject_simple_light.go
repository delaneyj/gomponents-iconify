package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 208a6 6 0 0 1-6 6H32a6 6 0 0 1 0-12h192a6 6 0 0 1 6 6ZM27.39 158a14 14 0 0 1 1.77-15l81.66-100.8a22.1 22.1 0 0 1 34.36 0L226.84 143a14.09 14.09 0 0 1-10.93 23H40.09a14 14 0 0 1-12.7-8Zm10.83-5.16a2 2 0 0 0 1.87 1.16h175.82a2 2 0 0 0 1.87-1.18a2 2 0 0 0-.27-2.24L135.86 49.76a10.1 10.1 0 0 0-15.72 0L38.49 150.58a2 2 0 0 0-.27 2.24Z"/>`),
		g.Group(children),
	)
}