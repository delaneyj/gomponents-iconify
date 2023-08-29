package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Surfshark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.15 18.47C38.7 5 39.15 5 27.73 4.55c-11-.45-13 2.17-15.07 11.66c-5.49 25.6-7.32 31.88 9.13 24.26c17.89-8.3 17.71-11.66 17.36-22Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.16 33.34h1A7.81 7.81 0 0 0 24 25.6v-5.72a7.81 7.81 0 0 1 7.86-7.74h1"/>`),
		g.Group(children),
	)
}