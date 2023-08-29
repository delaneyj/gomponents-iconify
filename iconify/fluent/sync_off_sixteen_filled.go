package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8A6 6 0 0 1 2 8h3.877l-.847 2.33a.5.5 0 0 0 .94.34l2-5.5a.5.5 0 1 0-.94-.34L6.24 7H2.084a6.002 6.002 0 0 1 11.834 0h-3.612l.665-1.83a.5.5 0 1 0-.94-.34l-2 5.5a.5.5 0 0 0 .94.34L9.941 8H14Z"/>`),
		g.Group(children),
	)
}