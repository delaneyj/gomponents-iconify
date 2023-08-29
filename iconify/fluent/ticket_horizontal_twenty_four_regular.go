package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketHorizontalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.25 5c.966 0 1.75.784 1.75 1.75v2.26c0 .39-.3.716-.688.748a2.25 2.25 0 0 0 0 4.484a.75.75 0 0 1 .688.748v2.26A1.75 1.75 0 0 1 20.25 19H3.75A1.75 1.75 0 0 1 2 17.25v-2.26c0-.39.3-.716.689-.748a2.25 2.25 0 0 0 0-4.484A.75.75 0 0 1 2 9.01V6.75C2 5.784 2.784 5 3.75 5h16.5Zm.25 3.385V6.75a.25.25 0 0 0-.25-.25H3.75a.25.25 0 0 0-.25.25v1.635a3.752 3.752 0 0 1 0 7.23v1.635c0 .138.112.25.25.25h16.5a.25.25 0 0 0 .25-.25v-1.635a3.752 3.752 0 0 1-.189-7.173l.19-.057Z"/>`),
		g.Group(children),
	)
}