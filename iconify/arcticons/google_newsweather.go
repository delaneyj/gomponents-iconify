package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleNewsweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 8a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h26.83v-9.17a2 2 0 0 1 2-2H44V10a2 2 0 0 0-2-2Zm38 20.83L32.83 40M9.08 33.42H26.5M9.08 14.58h29.84M9.08 24H26.5"/>`),
		g.Group(children),
	)
}