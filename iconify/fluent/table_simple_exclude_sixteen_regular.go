package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleExcludeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 1A2.5 2.5 0 0 0 1 3.5v6A2.5 2.5 0 0 0 3.5 12h3a.5.5 0 0 0 .5-.5V7h4.5a.5.5 0 0 0 .5-.5v-3A2.5 2.5 0 0 0 9.5 1h-6ZM6 7v4H3.5A1.5 1.5 0 0 1 2 9.5V7h4Zm0-1H2V3.5A1.5 1.5 0 0 1 3.5 2H6v4Zm1 0V2h2.5A1.5 1.5 0 0 1 11 3.5V6H7Zm1.5 4.25c0-.966.784-1.75 1.75-1.75h3c.966 0 1.75.784 1.75 1.75v3A1.75 1.75 0 0 1 13.25 15h-3a1.75 1.75 0 0 1-1.75-1.75v-3Z"/>`),
		g.Group(children),
	)
}