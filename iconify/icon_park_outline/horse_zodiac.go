package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorseZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M18 4C13 4 4 8 4 18.298V25m31 5v14M18.014 22c.486 1.5 1.986 4.5 6.382 3.738C26.98 25.427 33.082 26.14 36 31c1.5 2.5 5.447 2.496 8-3.86"/><path stroke-linejoin="round" d="m26 16l-2-2.5s-1.617-1.755-3-2.5c-1.383-.744-4-1.5-7 0c-1.297.815-4 2-4 6.943V44"/><path d="M27 44v-3a5 5 0 0 0-5-5v0a5 5 0 0 0-5 5v3"/></g>`),
		g.Group(children),
	)
}