package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Money(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 12h2m12 0h2M.5 16.5V20H1a43.131 43.131 0 0 1 3-.37M.5 16.5H1a3 3 0 0 1 3 3v.13M.5 16.5v-9M4 19.63a94.725 94.725 0 0 1 8-.38c2.134 0 5.281.127 8 .38m0 0a43.03 43.03 0 0 1 3 .37h.5v-3.5M20 19.63v-.13a3 3 0 0 1 3-3h.5m0 0v-9m0 0V4H23a43.03 43.03 0 0 1-3 .37m3.5 3.13H23a3 3 0 0 1-3-3v-.13m0 0a94.72 94.72 0 0 1-8 .38a94.72 94.72 0 0 1-8-.38m0 0A43.178 43.178 0 0 1 1 4H.5v3.5M4 4.37v.13a3 3 0 0 1-3 3H.5m11.5 7a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}