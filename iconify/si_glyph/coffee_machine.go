package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.021 15.085v.879H17v-2.928H3.027c-1.109 0-2.006.917-2.006 2.049zM6.035 12h.918c1.104 0 2-.926 2-2.067V8.037H4.035v.991h-1v.905h1c0 1.141.897 2.067 2 2.067z"/><path d="M1.083.083v1.834c0 1.076.851.947 1.917.991v1.05h2.042v2H6v.917h.959V3.958h1.916V2.917H11v9.067h6V.083H1.083zm4.959 3.001H4.917V1.916h1.125v1.168zm2 0H6.917V1.916h1.125v1.168z"/></g>`),
		g.Group(children),
	)
}