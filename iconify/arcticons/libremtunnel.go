package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libremtunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.77 37.7v5.8h-3.9v-5.78a13.88 13.88 0 0 0-27.76 0v5.78h-3.9V22.27a17.77 17.77 0 1 1 35.53 0Z"/>`),
		g.Group(children),
	)
}