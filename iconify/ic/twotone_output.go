package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneOutput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 17l5-5l-5-5l-1.41 1.41L18.17 11H9v2h9.17l-2.58 2.59z"/><path fill="currentColor" d="M19 19H5V5h14v2h2V5a2 2 0 0 0-2-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2v-2h-2v2z"/>`),
		g.Group(children),
	)
}