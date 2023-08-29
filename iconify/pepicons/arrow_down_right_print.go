package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><g opacity=".8"><path d="M16.143 6.856a1.5 1.5 0 0 1 1.358 1.63l-.472 5.185a1.5 1.5 0 0 1-2.987-.271l.471-5.186a1.5 1.5 0 0 1 1.63-1.358Z"/><path d="M8.856 14.143a1.5 1.5 0 0 1 1.358-1.63l5.186-.471a1.5 1.5 0 1 1 .271 2.987l-5.185.472a1.5 1.5 0 0 1-1.63-1.358Z"/><path d="M15.182 13.182a1.5 1.5 0 0 1-2.121 0L7.404 7.525a1.5 1.5 0 0 1 2.121-2.121l5.657 5.657a1.5 1.5 0 0 1 0 2.121Z"/></g><path d="M14.052 7.852a.5.5 0 0 1 .453.543l-.472 5.186a.5.5 0 1 1-.995-.09l.47-5.186a.5.5 0 0 1 .544-.453Z"/><path d="M7.852 14.052a.5.5 0 0 1 .453-.543l5.185-.471a.5.5 0 0 1 .09.995l-5.185.472a.5.5 0 0 1-.543-.453Z"/><path d="M13.354 13.354a.5.5 0 0 1-.708 0L6.111 6.818a.5.5 0 1 1 .707-.707l6.536 6.535a.5.5 0 0 1 0 .708Z"/></g>`),
		g.Group(children),
	)
}