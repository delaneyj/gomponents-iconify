package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodecademyGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.466h25.067v25.067H4.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.92 19.933a7.484 7.484 0 1 0 0 8.135m11.124 8.466H43.5"/>`),
		g.Group(children),
	)
}