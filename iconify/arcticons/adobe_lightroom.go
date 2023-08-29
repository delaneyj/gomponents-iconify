package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeLightroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.364 15v18h9m3.772-7.425a4.5 4.5 0 0 1 4.5-4.5h0m-4.5 0V33"/>`),
		g.Group(children),
	)
}