package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.576 1.334a.6.6 0 0 1 .849 0l2.121 2.121a.6.6 0 0 1-.848.849L10.6 3.207V7.5a.6.6 0 1 1-1.2 0V3.207L8.303 4.304a.6.6 0 0 1-.848-.849l2.121-2.121ZM3.9 10a.6.6 0 0 1 .6-.6h11a.6.6 0 1 1 0 1.2h-11a.6.6 0 0 1-.6-.6Zm6.7 6.794V12.5a.6.6 0 0 0-1.2 0v4.294l-1.097-1.097a.6.6 0 0 0-.848.848l2.121 2.122a.6.6 0 0 0 .849 0l2.121-2.122a.6.6 0 1 0-.848-.848L10.6 16.795Z"/>`),
		g.Group(children),
	)
}