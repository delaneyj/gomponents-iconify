package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fossilhybrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.33 20.3v7.41h3.7m-5.67-7.41v7.41M8.7 20.3h3.71M8.7 24h2.41M8.7 20.3v7.41m11.41-.81a2.07 2.07 0 0 0 1.82.81H23a1.85 1.85 0 0 0 1.85-1.85h0A1.85 1.85 0 0 0 23 24h-1.19A1.85 1.85 0 0 1 20 22.15h0a1.85 1.85 0 0 1 1.85-1.85h1.1a2.09 2.09 0 0 1 1.82.81m1.86 5.79a2.09 2.09 0 0 0 1.82.81h1.1a1.85 1.85 0 0 0 1.85-1.85h0A1.85 1.85 0 0 0 29.55 24h-1.22a1.85 1.85 0 0 1-1.85-1.85h0a1.85 1.85 0 0 1 1.85-1.85h1.1a2.09 2.09 0 0 1 1.82.81m-17.79 4.14a2.46 2.46 0 0 0 4.91 0v-2.51a2.46 2.46 0 0 0-4.91 0Z"/>`),
		g.Group(children),
	)
}