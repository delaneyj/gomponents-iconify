package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortDownLinesSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m9.03 11.03l-2.75 2.75a.75.75 0 0 1-1.06 0l-2.75-2.75a.75.75 0 1 1 1.06-1.06L5 11.44V2.75a.75.75 0 0 1 1.5 0v8.69l1.47-1.47a.75.75 0 0 1 1.06 1.06ZM8.5 2a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5h-5Zm0 2.5a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5h-3Zm0 2.5a.75.75 0 0 0 0 1.5h1a.75.75 0 0 0 0-1.5h-1Z"/>`),
		g.Group(children),
	)
}