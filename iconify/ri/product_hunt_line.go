package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductHuntLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 22c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm1.334-8a1.5 1.5 0 0 0 0-3H10.5v3h2.834Zm0-5a3.5 3.5 0 1 1 0 7H10.5v3h-2V7h4.834Z"/>`),
		g.Group(children),
	)
}