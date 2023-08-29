package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HsbcUk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.25 13.75h-20.5L24 24l10.25-10.25zm-20.5 20.5h20.5L24 24L13.75 34.25zm20.5-20.5v20.5L44.5 24L34.25 13.75zm-20.5 20.5v-20.5L3.5 24l10.25 10.25zm7.1 3v2.675a1.325 1.325 0 1 0 2.65 0V37.25m1.5 0v4m2.15 0l-1.645-2l1.645-1.986m-1.647 1.986H25"/>`),
		g.Group(children),
	)
}