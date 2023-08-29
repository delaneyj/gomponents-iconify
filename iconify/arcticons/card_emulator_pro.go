package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardEmulatorPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="29" x="4.5" y="9.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16.5h39m-10 18v-2m3 2v-2m3 2v-2"/>`),
		g.Group(children),
	)
}