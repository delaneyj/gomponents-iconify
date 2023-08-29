package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monolith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.581 43.5l-10.937-6.769V8.137l10.937 6.768V43.5zm-10.937-6.769l10.937-6.768"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.644 8.137L21.419 4.5l10.937 6.768v28.595L26.581 43.5m0-28.595l5.775-3.637"/>`),
		g.Group(children),
	)
}