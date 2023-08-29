package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.098 10.354a1 1 0 0 1 .905 1.087l-.472 5.185a1 1 0 1 1-1.991-.18l.471-5.186a1 1 0 0 1 1.087-.906Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.354 17.098a1 1 0 0 1 .906-1.087l5.185-.471a1 1 0 1 1 .181 1.991l-5.185.472a1 1 0 0 1-1.087-.905Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.828 15.829a1 1 0 0 1-1.414 0l-5.657-5.657a1 1 0 1 1 1.415-1.415l5.656 5.657a1 1 0 0 1 0 1.415Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}