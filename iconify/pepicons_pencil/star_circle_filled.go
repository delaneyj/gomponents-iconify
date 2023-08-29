package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilStarCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 5a.5.5 0 0 1 .435.253l2.425 4.274l4.745 1.023a.5.5 0 0 1 .27.82l-3.253 3.667l.51 4.911a.5.5 0 0 1-.703.507L13 18.448l-4.429 2.007a.5.5 0 0 1-.704-.507l.51-4.91l-3.251-3.668a.5.5 0 0 1 .269-.82l4.745-1.023l2.425-4.274A.5.5 0 0 1 13 5Zm0 1.513l-2.1 3.702a.5.5 0 0 1-.33.242l-4.128.89l2.83 3.191a.5.5 0 0 1 .123.384l-.443 4.263l3.842-1.74a.5.5 0 0 1 .412 0l3.842 1.74l-.443-4.263a.5.5 0 0 1 .123-.384l2.83-3.191l-4.128-.89a.5.5 0 0 1-.33-.242L13 6.513Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilStarCircleFilled0)"/></g>`),
		g.Group(children),
	)
}