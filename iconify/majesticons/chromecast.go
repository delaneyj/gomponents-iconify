package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chromecast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h.01M7 20a4 4 0 0 0-4-4m8 4a8 8 0 0 0-8-8"/><path fill="currentColor" d="M19 4H5a2 2 0 0 0-2 2v2.5c4 .167 12 2.7 12 11.5h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/></g>`),
		g.Group(children),
	)
}