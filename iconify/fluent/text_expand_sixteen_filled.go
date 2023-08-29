package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextExpandSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2.75A.75.75 0 0 1 2.75 2h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 2.75ZM4.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7ZM5 6.5a.5.5 0 0 0-1 0V8H2.5a.5.5 0 0 0 0 1H4v1.5a.5.5 0 0 0 1 0V9h1.5a.5.5 0 0 0 0-1H5V6.5Zm-2.25 7a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H2.75ZM9 10.75a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75ZM9.75 6a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Z"/>`),
		g.Group(children),
	)
}