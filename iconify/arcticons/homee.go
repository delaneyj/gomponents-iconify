package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Homee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.005 20.748c-1.826-6.973-7.328-10.02-12.617-9.736a12.025 12.025 0 0 0-10.127 6.475c-12.416 2.134-10.922 18.037 0 19.52h1.114v-5.6l7.252-5l7.253 5v5.6h7.185a8.159 8.159 0 0 0-.06-16.259ZM18.019 37.006v-2.417"/>`),
		g.Group(children),
	)
}