package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftPlanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.2 29.5h10.3c1.1 0 2-.9 2-2v-20c0-1.1-.9-2-2-2H30.2V34c0 1.1-.9 2-2 2H17.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 23.8h12.3v16.7c0 1.1-.9 2-2 2H7.5c-1.1 0-2-.9-2-2v-33c0-1.1.9-2 2-2h22.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.8 5.5v13.2h12.35"/>`),
		g.Group(children),
	)
}