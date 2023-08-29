package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrganizationTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.75 6.248a4.25 4.25 0 1 1 5 4.184V13.5h4.5a2.25 2.25 0 0 1 2.25 2.25v1.816a4.251 4.251 0 1 1-1.5 0V15.75a.75.75 0 0 0-.75-.75H8.75a.75.75 0 0 0-.75.75v1.816a4.251 4.251 0 1 1-1.5 0V15.75a2.25 2.25 0 0 1 2.25-2.25h4.5v-3.068a4.251 4.251 0 0 1-3.5-4.184Z"/>`),
		g.Group(children),
	)
}