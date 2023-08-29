package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magicplan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.5 42.5h8a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 30.956v-8.662c0-2.9-2.35-5.25-5.25-5.25h0a5.25 5.25 0 0 0-5.25 5.25m21 8.662v-8.662c0-2.9-2.35-5.25-5.25-5.25h0a5.25 5.25 0 0 0-5.25 5.25m-10.5 8.662V17.044"/>`),
		g.Group(children),
	)
}