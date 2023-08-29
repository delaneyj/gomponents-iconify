package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M13 5a.5.5 0 0 1 .435.253l2.425 4.274l4.745 1.023a.5.5 0 0 1 .27.82l-3.253 3.667l.51 4.911a.5.5 0 0 1-.703.507L13 18.448l-4.429 2.007a.5.5 0 0 1-.704-.507l.51-4.91l-3.251-3.668a.5.5 0 0 1 .269-.82l4.745-1.023l2.425-4.274A.5.5 0 0 1 13 5Zm0 1.513l-2.1 3.702a.5.5 0 0 1-.33.242l-4.128.89l2.83 3.191a.5.5 0 0 1 .123.384l-.443 4.263l3.842-1.74a.5.5 0 0 1 .412 0l3.842 1.74l-.443-4.263a.5.5 0 0 1 .123-.384l2.83-3.191l-4.128-.89a.5.5 0 0 1-.33-.242L13 6.513Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}