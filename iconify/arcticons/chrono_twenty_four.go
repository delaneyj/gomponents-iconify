package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChronoTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.203 39.203c-8.398 8.398-22.008 8.398-30.406 0S.4 17.195 8.797 8.797s22.008-8.398 30.406 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.405 11.405L24 24l8.814-8.814"/>`),
		g.Group(children),
	)
}