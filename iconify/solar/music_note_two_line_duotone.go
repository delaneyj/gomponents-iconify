package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M13 18V2" opacity=".5"/><circle cx="9" cy="18" r="4"/><path stroke-linecap="round" d="M19 8a6 6 0 0 1-6-6"/></g>`),
		g.Group(children),
	)
}