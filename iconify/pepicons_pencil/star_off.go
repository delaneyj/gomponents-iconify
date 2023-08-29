package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 2a.5.5 0 0 1 .435.253l2.425 4.274l4.745 1.023a.5.5 0 0 1 .27.82l-3.253 3.667l.51 4.911a.5.5 0 0 1-.703.507L10 15.448l-4.429 2.007a.5.5 0 0 1-.704-.507l.51-4.91l-3.25-3.668a.5.5 0 0 1 .269-.82L7.14 6.527l2.425-4.274A.5.5 0 0 1 10 2Zm0 1.513L7.9 7.215a.5.5 0 0 1-.33.242l-4.128.89l2.83 3.191a.5.5 0 0 1 .123.384l-.443 4.263l3.842-1.74a.5.5 0 0 1 .412 0l3.842 1.74l-.443-4.263a.5.5 0 0 1 .123-.384l2.83-3.191l-4.128-.89a.5.5 0 0 1-.33-.242L10 3.513Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}