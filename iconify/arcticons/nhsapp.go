package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nhsapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 15.79v16.421h39V15.79Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".999" d="M32.346 28.352a3.184 3.184 0 0 0 2.71 1.268h1.628a2.774 2.774 0 0 0 2.71-2.817h0a2.774 2.774 0 0 0-2.71-2.817H34.92a2.774 2.774 0 0 1-2.711-2.817h0a2.774 2.774 0 0 1 2.711-2.817h1.627a2.869 2.869 0 0 1 2.711 1.268"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.605 29.645v-11.29l7.185 11.29v-11.29m4.61 0v11.29M20.408 24h7.184m0-5.645v11.29"/>`),
		g.Group(children),
	)
}