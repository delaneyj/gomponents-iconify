package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaPlayReverseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 7v10l-5.1-5L14 7m-1.4-1.4C10 8.1 6 12 6 12s4 3.9 6.6 6.4c.4.4.9.6 1.4.6c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2c-.5 0-1 .2-1.4.6z"/>`),
		g.Group(children),
	)
}