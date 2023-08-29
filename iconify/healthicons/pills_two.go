package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 26c5.523 0 10-4.477 10-10S21.523 6 16 6S6 10.477 6 16s4.477 10 10 10Zm4.304-10.692a1 1 0 0 0-.762-1.849l-7.846 3.233a1 1 0 1 0 .763 1.85l7.845-3.234ZM32 42c5.523 0 10-4.477 10-10s-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10Zm3.91-8.075a1 1 0 1 0 .449-1.949l-8.27-1.901a1 1 0 1 0-.448 1.949l8.27 1.901Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}