package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractUpDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.793 5.207L12 11.414l6.207-6.207l-1.414-1.414L12 8.586L7.207 3.793L5.793 5.207Zm12.414 13.586L12 12.586l-6.207 6.207l1.414 1.414L12 15.414l4.793 4.793l1.414-1.414Z"/>`),
		g.Group(children),
	)
}