package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aimp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.11 33.07h0h-10.25L24 21.22l10.92 18.14h7.58L24 8.64L5.5 39.36h25.4l-3.79-6.29z"/>`),
		g.Group(children),
	)
}