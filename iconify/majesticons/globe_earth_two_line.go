package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeEarthTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M16 4v0a2 2 0 0 1-2 2h-.5A1.5 1.5 0 0 0 12 7.5V8a1 1 0 0 1-1 1v0a1 1 0 0 0-1 1v1a2 2 0 0 0 2 2h2a2 2 0 0 1 2 2v0a2 2 0 0 0 2 2h1m-8 4v-2.5c0-.828-.685-1.5-1.513-1.5v0C8.673 17 8 16.34 8 15.526v0c0-.34-.118-.67-.333-.933L3.5 9.5"/></g>`),
		g.Group(children),
	)
}