package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHorizontalTwoDashesSolidSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2ZM7 5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H7Zm4.5 0a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Zm-9 5a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1h-11Z"/>`),
		g.Group(children),
	)
}