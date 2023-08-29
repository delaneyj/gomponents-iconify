package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselVerticalFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18 6H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zm-2 13a2 2 0 0 1 1.995 1.85L18 21v1a1 1 0 0 1-1.993.117L16 22v-1H8v1a1 1 0 0 1-1.993.117L6 22v-1a2 2 0 0 1 1.85-1.995L8 19h8zm1-18a1 1 0 0 1 .993.883L18 2v1a2 2 0 0 1-1.85 1.995L16 5H8a2 2 0 0 1-1.995-1.85L6 3V2a1 1 0 0 1 1.993-.117L8 2v1h8V2a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}