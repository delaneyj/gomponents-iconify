package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10.063V4a1 1 0 0 0-1-1h-1c-1.979 1.979-5.697 3.087-8 3.613v10.774c2.303.526 6.021 1.634 8 3.613h1a1 1 0 0 0 1-1v-6.063a2 2 0 0 0 0-3.874ZM5 7a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h1l1 5h2V7H5Z"/>`),
		g.Group(children),
	)
}