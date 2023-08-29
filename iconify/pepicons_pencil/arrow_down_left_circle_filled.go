package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilArrowDownLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.148 17.052a.5.5 0 0 1-.543.453l-5.186-.471a.5.5 0 1 1 .09-.996l5.186.471a.5.5 0 0 1 .453.543Z"/><path d="M8.948 10.852a.5.5 0 0 1 .543.453l.471 5.185a.5.5 0 1 1-.995.09l-.472-5.185a.5.5 0 0 1 .453-.543Z"/><path d="M9.646 16.354a.5.5 0 0 1 0-.708l6.536-6.535a.5.5 0 0 1 .707.707l-6.535 6.536a.5.5 0 0 1-.708 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilArrowDownLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}