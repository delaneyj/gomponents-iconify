package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TachometerFastAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5a10 10 0 0 0-8.66 15a1 1 0 0 0 1.74-1A7.92 7.92 0 0 1 4 15a8 8 0 1 1 14.93 4a1 1 0 0 0 .37 1.37a1 1 0 0 0 1.36-.37A10 10 0 0 0 12 5Zm2.84 5.76l-1.55 1.54A2.91 2.91 0 0 0 12 12a3 3 0 1 0 3 3a2.9 2.9 0 0 0-.3-1.28l1.55-1.54a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0ZM12 16a1 1 0 0 1 0-2a1 1 0 0 1 .7.28a1 1 0 0 1 .3.72a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}