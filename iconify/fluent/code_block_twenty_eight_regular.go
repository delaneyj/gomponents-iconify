package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeBlockTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.75A3.75 3.75 0 0 1 6.75 3h14.5A3.75 3.75 0 0 1 25 6.75v14.5A3.75 3.75 0 0 1 21.25 25H6.75A3.75 3.75 0 0 1 3 21.25V6.75ZM6.75 4.5A2.25 2.25 0 0 0 4.5 6.75v14.5a2.25 2.25 0 0 0 2.25 2.25h14.5a2.25 2.25 0 0 0 2.25-2.25V6.75a2.25 2.25 0 0 0-2.25-2.25H6.75Zm5.28 4.72a.75.75 0 0 1 0 1.06L8.31 14l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Zm5 0a.75.75 0 1 0-1.06 1.06L19.69 14l-3.72 3.72a.75.75 0 1 0 1.06 1.06l4.25-4.25a.75.75 0 0 0 0-1.06l-4.25-4.25Z"/>`),
		g.Group(children),
	)
}