package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coinstats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.692 30.754A13.5 13.5 0 1 1 24 10.5M12.31 30.745l-6.93 4.001"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.617 34.752L24.003 24.006V2.501"/>`),
		g.Group(children),
	)
}