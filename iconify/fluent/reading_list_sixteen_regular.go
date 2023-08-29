package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadingListSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4.999a1.001 1.001 0 0 1 1.835-.555a.5.5 0 1 0 .832-.555A2 2 0 0 0 2 4.999A2.001 2.001 0 0 0 4.077 7h7.426a.5.5 0 0 0 0-1H4.077a.51.51 0 0 0-.038.001A1.001 1.001 0 0 1 3 4.998ZM7.499 4a.5.5 0 0 0 0 1h6a.5.5 0 1 0 0-1h-6ZM4.495 8.002a.5.5 0 0 0 0 1h9a.5.5 0 1 0 0-1h-9ZM2 10.5a.5.5 0 0 1 .5-.5h8.998a.5.5 0 1 1 0 1H2.5a.5.5 0 0 1-.5-.5ZM4.495 12a.5.5 0 0 0 0 1h9a.5.5 0 1 0 0-1h-9Z"/>`),
		g.Group(children),
	)
}