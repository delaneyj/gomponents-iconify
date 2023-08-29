package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFiveLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174 160a54 54 0 0 1-90 40.25a6 6 0 1 1 8-8.94a42 42 0 1 0 .8-63.31a6 6 0 0 1-9.8-5.75l15.12-75.43A6 6 0 0 1 104 42h64a6 6 0 0 1 0 12h-59.08l-11.38 56.89A54 54 0 0 1 174 160Z"/>`),
		g.Group(children),
	)
}