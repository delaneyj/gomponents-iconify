package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUpSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 8a5 5 0 1 1 10 0A5 5 0 0 1 3 8Zm5-6a6 6 0 1 0 0 12A6 6 0 0 0 8 2Zm2.854 6.646l-2.5-2.5a.5.5 0 0 0-.708 0l-2.5 2.5a.5.5 0 1 0 .708.708L8 7.207l2.146 2.147a.5.5 0 0 0 .708-.708Z"/>`),
		g.Group(children),
	)
}