package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReadMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.24 1.061a.5.5 0 0 0-.48 0L1.7 3.821c-.202.11-.367.267-.487.453L7 7.43l5.787-3.156a1.34 1.34 0 0 0-.487-.452L7.24 1.06ZM1 10V5.297l5.76 3.142a.5.5 0 0 0 .48 0L13 5.297V10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Zm2.268 3A2 2 0 0 0 5 14h6a4 4 0 0 0 4-4V6a2 2 0 0 0-1-1.732V10a3 3 0 0 1-3 3H3.268Z"/>`),
		g.Group(children),
	)
}