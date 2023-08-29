package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypePattern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15H17V2h-2v13H2v2h13v13h2V17h13v-2z"/><path fill="currentColor" d="M25.586 20L27 21.414L23.414 25L27 28.586L25.586 30l-5-5l5-5zM11 30H3a1 1 0 0 1-.894-1.447l4-8a1.041 1.041 0 0 1 1.789 0l4 8A1 1 0 0 1 11 30zm-6.382-2h4.764L7 23.236zM28 12h-6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zm-6-8v6h6.001L28 4zM7 12a5 5 0 1 1 5-5a5.006 5.006 0 0 1-5 5zm0-8a3 3 0 1 0 3 3a3.003 3.003 0 0 0-3-3z"/>`),
		g.Group(children),
	)
}