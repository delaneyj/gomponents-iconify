package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoPersonClockSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM1.75 1A1.75 1.75 0 0 0 0 2.75v6.5C0 10.216.784 11 1.75 11h4.272a5.475 5.475 0 0 1 1.235-3H5a1 1 0 0 0-1 1v1H1.75A.75.75 0 0 1 1 9.25v-6.5A.75.75 0 0 1 1.75 2h10.5a.75.75 0 0 1 .75.75v3.457c.349.099.683.23 1 .393V2.75A1.75 1.75 0 0 0 12.25 1H1.75ZM16 11.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0ZM11.5 9a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5H13a.5.5 0 0 0 0-1h-1V9.5a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}