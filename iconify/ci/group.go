package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Group(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3a5 5 0 1 0 0 10A5 5 0 0 0 9 3ZM6 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm10.908.218A2 2 0 0 0 16 8V6a4 4 0 1 1-2.357 7.232l1.178-1.616a2 2 0 1 0 2.087-3.398ZM19.998 21A3.999 3.999 0 0 0 16 17.002V15a6.001 6.001 0 0 1 6 6h-2.002ZM16 21h-2a5 5 0 0 0-10 0H2a7 7 0 1 1 14 0Z"/>`),
		g.Group(children),
	)
}