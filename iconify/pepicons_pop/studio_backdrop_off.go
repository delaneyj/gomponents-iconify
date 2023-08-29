package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 3a1.5 1.5 0 0 1 1.5-1.5h11A1.5 1.5 0 0 1 17 3v7a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 3 10V3Zm2 .5v6h10v-6H5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M1.75 2.5a1 1 0 0 1 1-1h14.5a1 1 0 1 1 0 2H2.75a1 1 0 0 1-1-1Zm1.97 14l1.26-6.304l-1.96-.392l-1.381 6.902a1.5 1.5 0 0 0 1.47 1.794H16.89a1.5 1.5 0 0 0 1.471-1.794l-1.38-6.902l-1.962.392L16.28 16.5H3.72Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}