package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 26H4a2.002 2.002 0 0 1-2-2v-4h2v4h4zM2 12h2v4H2zm24-4h-2V4h-4V2h4a2.002 2.002 0 0 1 2 2zM12 2h4v2h-4zM4 8H2V4a2.002 2.002 0 0 1 2-2h4v2H4zm23 24a.997.997 0 0 1-.707-.293l-6.138-6.138l-3.323 4.986a1 1 0 0 1-1.79-.268l-6-20a1 1 0 0 1 1.245-1.245l20 6a1 1 0 0 1 .268 1.79l-4.986 3.323l6.138 6.138a1 1 0 0 1 0 1.414l-4 4A.997.997 0 0 1 27 32zm0-2.414L29.586 27l-7.155-7.155l5.246-3.498l-16.185-4.855l4.855 16.185l3.498-5.246z"/>`),
		g.Group(children),
	)
}