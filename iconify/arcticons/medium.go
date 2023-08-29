package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2m0-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2"/><circle cx="15.329" cy="24" r="9.068" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="31.198" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.269" ry="9.068"/><ellipse cx="39.869" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.87" ry="9.068"/>`),
		g.Group(children),
	)
}