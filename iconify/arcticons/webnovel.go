package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webnovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.674 27.544c-1.675-.822-2.897-2.471-2.913-4.337v-.044c0-3.797 4.04-6.65 8.055-4.497c1.23.659 2.04 1.906 2.29 3.278h0a5.229 5.229 0 0 1-2.84 5.63L4.5 38.26V9.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 9.739v28.522L24 28.686"/>`),
		g.Group(children),
	)
}