package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 6a1.5 1.5 0 0 1 1.5-1.5h11A1.5 1.5 0 0 1 20 6v7a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 6 13V6Zm2 .5v6h10v-6H8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.75 5.5a1 1 0 0 1 1-1h14.5a1 1 0 1 1 0 2H5.75a1 1 0 0 1-1-1Zm1.97 14l1.26-6.304l-1.96-.392l-1.381 6.902a1.5 1.5 0 0 0 1.47 1.794H19.89a1.5 1.5 0 0 0 1.471-1.794l-1.38-6.902l-1.962.392L19.28 19.5H6.72Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}