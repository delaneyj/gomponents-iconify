package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractLeftRightLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.793 5.793L12.586 12l6.207 6.207l1.415-1.414L15.415 12l4.793-4.793l-1.415-1.414ZM5.207 18.207L11.414 12L5.207 5.793L3.793 7.207L8.586 12l-4.793 4.793l1.414 1.414Z"/>`),
		g.Group(children),
	)
}