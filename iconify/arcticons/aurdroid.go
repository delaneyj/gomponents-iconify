package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aurdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 34.21l-20.33 5.72V19.08l20.33-5.71v20.84zm-39 .76l18.67 4.96V19.08L4.5 14.12v20.85z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 14.12L24 8.07l19.5 5.3m-7.9 2.22v7.99l-5.08 1.56v-8.12m2.54-.72l-17.7-5.2"/>`),
		g.Group(children),
	)
}