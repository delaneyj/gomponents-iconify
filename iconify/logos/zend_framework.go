package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZendFramework(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="#6FB643"><path d="M120.43 36.02s0-33.818 33.82-33.818h101.456s0 33.819-33.82 33.819H120.43m0 43.212s0-33.819 33.82-33.819h50.727s0 33.82-33.82 33.82H120.43m0 43.212s0-33.818 33.82-33.818h16.907s0 33.818-33.817 33.818h-16.91"/><path d="M109.076 96.055v26.98H.32l70.06-95.568H10.236V.486h113.456L53.956 96.055h55.12"/></g>`),
		g.Group(children),
	)
}