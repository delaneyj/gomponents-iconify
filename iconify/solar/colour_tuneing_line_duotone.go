package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColourTuneingLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2 12h7.5M22 12h-7.5" opacity=".5"/><path d="M20 15.684C20 19 17.735 22 16 22c-2.268 0-3.928-3.158-3.928-10c0-6.842-1.66-10-3.928-10c-1.734 0-4 3-4 6.316"/></g>`),
		g.Group(children),
	)
}