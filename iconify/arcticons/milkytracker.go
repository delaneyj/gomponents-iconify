package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milkytracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.66 20.54v16l13.3 7l13.38-7V21l-6.25-8l-13.65-6.6ZM24 43.5v-16"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.66 20.54L24 27.47L37.34 21M24 27.47L31.09 13m6.25 8l-7.62-5.22M31.09 13v-1.8L17.44 4.5v1.9"/>`),
		g.Group(children),
	)
}