package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessKnightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m8.959 1.99l-.147.028l-.115.029a1 1 0 0 0-.646 1.27L8.8 5.562L5.985 7.297a2 2 0 0 0-.655 2.751l.089.133A2 2 0 0 0 7.033 11l1.563-.001l-1.614 4.674A1 1 0 0 0 7.927 17h7.961a1 1 0 0 0 1-.978l.112-5c0-3.827-1.555-6.878-4.67-7.966l-2.399-.83l-.375-.121l-.258-.074L9.163 2l-.101-.013l-.055-.001l-.048.003zM18 18H6a1 1 0 0 0-1 1a2 2 0 0 0 2 2h10a2 2 0 0 0 1.987-1.768l.011-.174A1 1 0 0 0 18 18z"/></g>`),
		g.Group(children),
	)
}