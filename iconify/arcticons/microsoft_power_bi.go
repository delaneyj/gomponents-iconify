package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftPowerBi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.379 23.613h9.508a2 2 0 0 1 2 2v18.036H11.379a2 2 0 0 1-2-2V25.613a2 2 0 0 1 2-2Zm15.58-9.44V6.649a2 2 0 0 1 2-2h9.508a2 2 0 0 1 2 2v35a2 2 0 0 1-2 2h-6.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.169 23.613v-7.44a2 2 0 0 1 2-2h9.508a2 2 0 0 1 2 2v27.476h-8.79"/>`),
		g.Group(children),
	)
}