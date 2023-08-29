package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFitInTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.21 7.781A.75.75 0 1 1 6.27 6.72l2.51 2.5a.75.75 0 0 1 0 1.061l-2.5 2.5a.75.75 0 0 1-1.06-1.06l1.22-1.22H2.75a.75.75 0 0 1 0-1.5h3.684l-1.223-1.22Zm9.58 0a.75.75 0 0 0-1.06-1.062l-2.51 2.5a.75.75 0 0 0 0 1.061l2.5 2.5a.75.75 0 1 0 1.06-1.06l-1.22-1.22h3.69a.75.75 0 0 0 0-1.5h-3.684l1.223-1.219Z"/>`),
		g.Group(children),
	)
}