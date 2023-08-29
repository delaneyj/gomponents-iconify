package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyAudi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9.567" cy="24" r="6.067"/><circle cx="19.189" cy="24" r="6.067"/><circle cx="28.811" cy="24" r="6.067"/><circle cx="38.433" cy="24" r="6.067"/></g>`),
		g.Group(children),
	)
}