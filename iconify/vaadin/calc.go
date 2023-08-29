package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 3h6v2H9V3zm0 8h6v2H9v-2zM5 1H3v2H1v2h2v2h2V5h2V3H5zm2 9.4L5.6 9L4 10.6L2.4 9L1 10.4L2.6 12L1 13.6L2.4 15L4 13.4L5.6 15L7 13.6L5.4 12zm6 4.1a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm0-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}