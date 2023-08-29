package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnreadSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.64 5.396A2 2 0 1 0 11.062 3a2 2 0 0 0 2.576 2.396ZM4 3h6.041a3.001 3.001 0 0 0 1.902 3.309L8 8.432l-6-3.23V5a2 2 0 0 1 2-2Zm10 3.337L8.237 9.44a.5.5 0 0 1-.474 0L2 6.337V11a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V6.337Z"/>`),
		g.Group(children),
	)
}