package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnreadSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 5.232A2 2 0 1 0 11.063 3a2.005 2.005 0 0 0 0 1A2 2 0 0 0 14 5.232ZM4 3h6.041a3.018 3.018 0 0 0 0 1H4a1 1 0 0 0-1 1v.74l5 2.692l3.944-2.123a2.993 2.993 0 0 0 2.056.02V11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2ZM3 6.876V11a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V6.876L8.237 9.44a.5.5 0 0 1-.474 0L3 6.876Z"/>`),
		g.Group(children),
	)
}