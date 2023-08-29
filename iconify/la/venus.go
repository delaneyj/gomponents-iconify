package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Venus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3c-4.406 0-8 3.594-8 8c0 4.066 3.066 7.438 7 7.938V23h-4v2h4v4h2v-4h4v-2h-4v-4.063c3.934-.5 7-3.87 7-7.937c0-4.406-3.594-8-8-8zm0 2c3.324 0 6 2.676 6 6s-2.676 6-6 6s-6-2.676-6-6s2.676-6 6-6z"/>`),
		g.Group(children),
	)
}