package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnreadTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM3.5 3h4.063a2.004 2.004 0 0 0 .593 1.981L6 5.951L2.037 4.169A1.5 1.5 0 0 1 3.5 3Zm2.705 3.956l3.237-1.457A2.01 2.01 0 0 0 10 5.437V7.5A1.5 1.5 0 0 1 8.5 9h-5A1.5 1.5 0 0 1 2 7.5V5.248l3.794 1.708a.5.5 0 0 0 .41 0Z"/>`),
		g.Group(children),
	)
}