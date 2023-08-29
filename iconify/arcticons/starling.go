package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Starling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 24v-1.5a18 18 0 0 1 18-18h1.5m0 19.5v1.5a18 18 0 0 1-18 18h-1.5"/>`),
		g.Group(children),
	)
}