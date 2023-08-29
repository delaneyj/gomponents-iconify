package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppMarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 28a3 3 0 1 0-3 3h3v-3Zm3 3a3 3 0 1 0-3-3v3h3Zm-3 3a3 3 0 1 0 3-3h-3v3Zm-3-3a3 3 0 1 0 3 3v-3h-3Zm-5-19.564a8 8 0 1 0 16 0"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}