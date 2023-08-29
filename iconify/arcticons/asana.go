package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="13.3" cy="33.42" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.8" ry="7.71"/><ellipse cx="34.7" cy="33.42" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.8" ry="7.71"/><ellipse cx="24" cy="14.58" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.8" ry="7.71"/>`),
		g.Group(children),
	)
}