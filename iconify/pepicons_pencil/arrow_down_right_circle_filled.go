package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilArrowDownRightCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.052 10.852a.5.5 0 0 1 .453.543l-.472 5.186a.5.5 0 1 1-.995-.09l.47-5.186a.5.5 0 0 1 .544-.453Z"/><path d="M10.852 17.052a.5.5 0 0 1 .453-.543l5.185-.471a.5.5 0 0 1 .09.995l-5.185.472a.5.5 0 0 1-.543-.453Z"/><path d="M16.354 16.354a.5.5 0 0 1-.708 0L9.111 9.818a.5.5 0 1 1 .707-.707l6.536 6.535a.5.5 0 0 1 0 .708Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilArrowDownRightCircleFilled0)"/></g>`),
		g.Group(children),
	)
}