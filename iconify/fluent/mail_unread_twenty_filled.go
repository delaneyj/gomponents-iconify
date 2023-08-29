package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnreadTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 6a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-13-3h9.535a3.499 3.499 0 0 0 1.72 3.535L10 9.92L2.015 5.223A2.5 2.5 0 0 1 4.5 3Zm5.754 7.931l6.742-3.967a3.53 3.53 0 0 0 1.004 0V13.5a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 2 13.5V6.373l7.747 4.558a.5.5 0 0 0 .507 0Z"/>`),
		g.Group(children),
	)
}