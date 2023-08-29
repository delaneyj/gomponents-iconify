package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FossMoodle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.5 18.874L23.936 7.002L3.5 18.8l20.564 11.873L44.5 18.874z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.065 23.272v10.235l12.975 7.491l12.895-7.444V23.272m7.565-4.398v11.233"/>`),
		g.Group(children),
	)
}