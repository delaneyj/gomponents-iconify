package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixAddSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 1.5a.5.5 0 0 1 .5-.5H8c.5 0 .989.053 1.46.152a7.01 7.01 0 0 1 5.504 6.132A6.98 6.98 0 0 1 12.899 13h-1.581A6 6 0 0 0 8 2H1.5a.5.5 0 0 1-.5-.5ZM1 8a7 7 0 0 0 7 7h6.5a.5.5 0 0 0 0-1H8a6.008 6.008 0 0 1-6-6a5.995 5.995 0 0 1 2.682-5H3.101A6.979 6.979 0 0 0 1 8Zm7-3a.5.5 0 0 1 .5.5v2h2a.5.5 0 0 1 0 1h-2v2a.5.5 0 0 1-1 0v-2h-2a.5.5 0 0 1 0-1h2v-2A.5.5 0 0 1 8 5Z"/>`),
		g.Group(children),
	)
}