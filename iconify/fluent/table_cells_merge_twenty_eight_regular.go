package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsMergeTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75v14.5A3.75 3.75 0 0 0 6.75 25h14.5A3.75 3.75 0 0 0 25 21.25V6.75A3.75 3.75 0 0 0 21.25 3H6.75Zm14.5 1.5a2.25 2.25 0 0 1 2.25 2.25v.75h-19v-.75A2.25 2.25 0 0 1 6.75 4.5h14.5ZM4.5 9h19v10h-19V9Zm0 12.25v-.75h19v.75a2.25 2.25 0 0 1-2.25 2.25H6.75a2.25 2.25 0 0 1-2.25-2.25Zm14.08-8H9.42l.89-1.002a.75.75 0 0 0-1.12-.996l-2 2.25a.75.75 0 0 0 0 .996l2 2.25a.75.75 0 1 0 1.12-.996l-.89-1.002h9.16l-.89 1.002a.75.75 0 0 0 1.12.996l2-2.25l.011-.012a.746.746 0 0 0-.013-.987l-1.997-2.247a.75.75 0 0 0-1.122.996l.89 1.002Z"/>`),
		g.Group(children),
	)
}