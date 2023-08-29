package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qtranslate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.874 42.432l7.639-22.949M42.835 42.5l-7.322-23.017M40.385 34.8h-9.971M5.165 10.582h25.653M17.991 5.5v5.082m7.179 0c0 7.018-8.147 18.634-16.294 20.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.167 16.955c.806 4.84 9.196 12.907 16.658 14.036"/>`),
		g.Group(children),
	)
}