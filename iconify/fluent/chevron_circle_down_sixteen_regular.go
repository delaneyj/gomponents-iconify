package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleDownSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 8a5 5 0 1 0 10 0A5 5 0 0 0 3 8Zm5 6A6 6 0 1 1 8 2a6 6 0 0 1 0 12Zm2.854-6.646l-2.5 2.5a.5.5 0 0 1-.708 0l-2.5-2.5a.5.5 0 1 1 .708-.708L8 8.793l2.146-2.147a.5.5 0 0 1 .708.708Z"/>`),
		g.Group(children),
	)
}