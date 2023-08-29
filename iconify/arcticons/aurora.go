package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aurora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.33 37H37.6a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2H12.73"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.465 30.5l-5.2-16l-5.4 16m1.8-5.4h7"/>`),
		g.Group(children),
	)
}