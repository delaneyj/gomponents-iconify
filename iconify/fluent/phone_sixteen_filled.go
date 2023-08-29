package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2.75C4 1.784 4.784 1 5.75 1h4.5c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0 1 10.25 15h-4.5A1.75 1.75 0 0 1 4 13.25V2.75ZM7 12a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H7Z"/>`),
		g.Group(children),
	)
}