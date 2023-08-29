package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buttercup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.616 16.09a11.59 11.59 0 1 1 23.179 0m0 4.23H6.615V43.5h23.18m0-23.18a11.59 11.59 0 0 1 11.59 11.59h0a11.59 11.59 0 0 1-11.59 11.59h0M6.616 20.32v-4.23m23.179 4.23v-4.23"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.038 32.492a3.16 3.16 0 1 1 2.162 0m-2.162 0l-.008 3.394m2.17-3.394v3.394m0 0a1.085 1.085 0 0 1-2.17 0"/>`),
		g.Group(children),
	)
}