package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirportLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15h-6.07A8.008 8.008 0 0 0 17 8.07V2h-2v6.07A8.008 8.008 0 0 0 8.07 15H2v2h6.07A8.008 8.008 0 0 0 15 23.93V30h2v-6.07A8.008 8.008 0 0 0 23.93 17H30Zm-14 7a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Z"/>`),
		g.Group(children),
	)
}