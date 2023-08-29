package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiomack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 23.432h1.274m39.726.942c-2.77-.083-3.85 1.911-4.793 3.99l-1.8-.97l.97-16.096L29.9 36.702l.11-21.054s-5.706 7.867-7.48 12.161c-1.883-3.657-2.548-7.618-2.548-7.618l-5.43 6.953l-1.156-5.236l-3.442 2.715s-1.122.042-2.968-1.19"/>`),
		g.Group(children),
	)
}