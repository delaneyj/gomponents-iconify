package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackLargeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#354a54" d="M63.998 57.1a6.899 6.899 0 0 1-6.899 6.903h-50.2A6.9 6.9 0 0 1 .003 57.1V6.9C.003 3.09 3.091 0 6.899 0h50.2a6.898 6.898 0 0 1 6.899 6.9v50.2"/>`),
		g.Group(children),
	)
}