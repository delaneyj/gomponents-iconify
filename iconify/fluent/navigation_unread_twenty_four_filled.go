package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationUnreadTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.25 8.5a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5ZM15.713 7a3.744 3.744 0 0 1-.138-2H3l-.117.007A1 1 0 0 0 3 7h12.713ZM21 17H3l-.117.007A1 1 0 0 0 3 19h18l.117-.007A1 1 0 0 0 21 17ZM3 11l18-.002a1 1 0 0 1 .117 1.993l-.117.007L3 13a1 1 0 0 1-.117-1.993L3 11Z"/>`),
		g.Group(children),
	)
}