package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.56 12.098l1.532 2.652A3.5 3.5 0 0 1 18.061 20h-2.063v2.5l-5-3.5l5-3.5V18h2.063a1.5 1.5 0 0 0 1.299-2.25l-1.532-2.652l1.733-1ZM7.304 9.133l.53 6.08l-2.165-1.25l-1.03 1.786A1.5 1.5 0 0 0 5.937 18h3.062v2H5.937a3.5 3.5 0 0 1-3.032-5.25l1.031-1.787l-2.165-1.249l5.532-2.58Zm6.446-6.165a3.5 3.5 0 0 1 1.28 1.281l1.032 1.786l2.165-1.25l-.531 6.08l-5.531-2.58l2.165-1.25l-1.031-1.786a1.5 1.5 0 0 0-2.598 0L9.168 7.903l-1.732-1L8.968 4.25a3.5 3.5 0 0 1 4.78-1.281Z"/>`),
		g.Group(children),
	)
}