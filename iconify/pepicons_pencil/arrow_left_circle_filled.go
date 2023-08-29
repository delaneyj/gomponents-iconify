package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilArrowLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.653 17.384a.5.5 0 0 1-.704-.064l-3.333-4a.5.5 0 0 1 .768-.64l3.333 4a.5.5 0 0 1-.064.704Z"/><path d="M11.653 8.616a.5.5 0 0 1 .064.704l-3.333 4a.5.5 0 0 1-.768-.64l3.333-4a.5.5 0 0 1 .704-.064Z"/><path d="M8 13a.5.5 0 0 1 .5-.5H18a.5.5 0 0 1 0 1H8.5A.5.5 0 0 1 8 13Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilArrowLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}