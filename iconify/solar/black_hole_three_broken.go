package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackHoleThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="2"/><path stroke-linecap="round" d="M12 10c5 0 4.6 12-3 12"/><path stroke-linecap="round" d="M12.312 14c-5 0-4.6-12 3-12"/><path stroke-linecap="round" d="M10 12.312c0-2.78 3.707-3.89 7-3.024m5 6.024c0-1.97-.806-3.456-2-4.49M14 12c0 2.779-3.707 3.89-7 3.024M2 9c0 1.68.586 3.008 1.5 4.004"/></g>`),
		g.Group(children),
	)
}