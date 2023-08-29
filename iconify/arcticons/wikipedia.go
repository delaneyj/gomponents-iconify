package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wikipedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.07h7.3m4.99 0h5.52m21.19 0h-6.58m-4.84 0h-4.83m-7.7 0l10.53 25.86l10.11-25.86"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.67 11.07L18.23 36.93L8.15 11.07"/>`),
		g.Group(children),
	)
}