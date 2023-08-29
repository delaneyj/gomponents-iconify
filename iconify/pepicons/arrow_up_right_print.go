package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><g opacity=".8"><path d="M8.856 6.357A1.5 1.5 0 0 1 10.486 5l5.185.472a1.5 1.5 0 0 1-.271 2.987l-5.186-.471a1.5 1.5 0 0 1-1.358-1.63Z"/><path d="M16.143 13.644a1.5 1.5 0 0 1-1.63-1.358L14.042 7.1a1.5 1.5 0 0 1 2.987-.271l.472 5.185a1.5 1.5 0 0 1-1.358 1.63Z"/><path d="M15.182 7.318a1.5 1.5 0 0 1 0 2.121l-5.657 5.657a1.5 1.5 0 1 1-2.121-2.121l5.657-5.657a1.5 1.5 0 0 1 2.12 0Z"/></g><path d="M7.852 5.952a.5.5 0 0 1 .543-.453l5.186.472a.5.5 0 0 1-.09.996l-5.186-.472a.5.5 0 0 1-.453-.543Z"/><path d="M14.052 12.152a.5.5 0 0 1-.543-.453l-.472-5.185a.5.5 0 0 1 .996-.09l.472 5.185a.5.5 0 0 1-.453.543Z"/><path d="M13.354 6.65a.5.5 0 0 1 0 .708l-6.536 6.535a.5.5 0 0 1-.707-.707l6.535-6.536a.5.5 0 0 1 .707 0Z"/></g>`),
		g.Group(children),
	)
}