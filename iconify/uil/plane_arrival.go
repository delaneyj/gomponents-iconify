package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaneArrival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.12 16.23a5 5 0 0 0-2.3-3L16.71 12l-.48-5.47a1 1 0 0 0-.49-.78l-3-1.72a1 1 0 0 0-1 0a1 1 0 0 0-.52.84l-.15 3.9l-1.75-1l-2.86-3.85a1 1 0 0 0-1.78.41l-.87 4.61a3 3 0 0 0 1.39 3.29l14.06 8.11a1 1 0 0 0 1.36-.34a4.91 4.91 0 0 0 .5-3.77ZM19.24 18L6.2 10.5a1 1 0 0 1-.44-1.13l.46-2.44l1.66 2.2a1 1 0 0 0 .3.27l3.35 1.94a1 1 0 0 0 1.5-.83l.16-3.9l1.09.63l.48 5.47a1 1 0 0 0 .5.78L17.82 15a2.91 2.91 0 0 1 1.36 1.78a2.74 2.74 0 0 1 .06 1.22Z"/>`),
		g.Group(children),
	)
}