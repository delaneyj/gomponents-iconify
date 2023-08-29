package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardFifteenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10h-2a8 8 0 1 1-1.384-4.5H16v1h-4v4.25h2.875a.625.625 0 1 1 0 1.25H12v1.5h2.875a2.125 2.125 0 0 0 0-4.25H13.5V10h3.25v-.5H22v-6h-2V6a9.985 9.985 0 0 0-8-4ZM8.5 8.5H10v7H8.5v-7Z"/>`),
		g.Group(children),
	)
}