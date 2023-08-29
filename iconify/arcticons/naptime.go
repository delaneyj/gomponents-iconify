package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Naptime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.474 9.641H43.5L24.474 38.359H43.5m-39-19.627h14.774L4.5 38.307h14.774"/>`),
		g.Group(children),
	)
}