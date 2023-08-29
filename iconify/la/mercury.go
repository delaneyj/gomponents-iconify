package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mercury(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 2c0 1.477.66 2.8 1.688 3.719C9.93 6.984 8 9.777 8 13c0 4.066 3.066 7.438 7 7.938V24h-4v2h4v4h2v-4h4v-2h-4v-3.063c3.934-.5 7-3.87 7-7.937c0-3.223-1.93-6.016-4.688-7.281C20.34 4.8 21 3.477 21 2h-2c0 1.668-1.332 3-3 3s-3-1.332-3-3zm5 5c3.324 0 6 2.676 6 6s-2.676 6-6 6s-6-2.676-6-6s2.676-6 6-6z"/>`),
		g.Group(children),
	)
}