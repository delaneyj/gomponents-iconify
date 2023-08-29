package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnsTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 3h8.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H12m0-18H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H12m0-18v18"/>`),
		g.Group(children),
	)
}