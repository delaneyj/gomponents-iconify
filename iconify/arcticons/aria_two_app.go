package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AriaTwoApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.116 26.932h17.768m4.616-8.329L23.999 43.5L10.5 18.603M24 21.427V4.5"/>`),
		g.Group(children),
	)
}