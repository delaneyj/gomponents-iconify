package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yeelight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.703 30.995L4.5 26.62V13.764L24 21.38v12.856m8.297-3.241L43.5 26.62V13.764L24 21.38v12.856"/>`),
		g.Group(children),
	)
}