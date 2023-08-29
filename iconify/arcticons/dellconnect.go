package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dellconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.11 11.93v-5a2.43 2.43 0 0 1 2.43-2.43h20.24a2.43 2.43 0 0 1 2.43 2.44v34.12a2.43 2.43 0 0 1-2.43 2.44H17.54a2.43 2.43 0 0 1-2.43-2.44v-9.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.93 27.25h22.96l-7.96 7.96M31.89 24H8.93l7.97-7.97"/>`),
		g.Group(children),
	)
}