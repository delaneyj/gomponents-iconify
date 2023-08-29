package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.052 7.852a.5.5 0 0 1 .453.543l-.472 5.186a.5.5 0 1 1-.995-.09l.47-5.186a.5.5 0 0 1 .544-.453Z"/><path d="M7.852 14.052a.5.5 0 0 1 .453-.543l5.185-.471a.5.5 0 0 1 .09.995l-5.185.472a.5.5 0 0 1-.543-.453Z"/><path d="M13.354 13.354a.5.5 0 0 1-.708 0L6.111 6.818a.5.5 0 1 1 .707-.707l6.536 6.535a.5.5 0 0 1 0 .708Z"/></g>`),
		g.Group(children),
	)
}