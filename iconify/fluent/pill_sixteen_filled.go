package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3.05A3.5 3.5 0 1 1 12.95 8L8 12.95A3.5 3.5 0 0 1 3.05 8L8 3.05Zm2.122 6.364l2.12-2.12a2.5 2.5 0 0 0-3.535-3.536l-2.121 2.12l3.536 3.536Zm-2.268 1.44a.5.5 0 1 0-.708-.707l-1.5 1.5a.5.5 0 1 0 .708.707l1.5-1.5Z"/>`),
		g.Group(children),
	)
}