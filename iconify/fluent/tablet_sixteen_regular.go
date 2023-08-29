package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 10a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3ZM2.75 3A1.75 1.75 0 0 0 1 4.75v6.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0 0 15 11.25v-6.5A1.75 1.75 0 0 0 13.25 3H2.75ZM2 4.75A.75.75 0 0 1 2.75 4h10.5a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-.75.75H2.75a.75.75 0 0 1-.75-.75v-6.5Z"/>`),
		g.Group(children),
	)
}