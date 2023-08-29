package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.25 18.392a2 2 0 0 1 0-2.828L15.564 4.25a2 2 0 0 1 2.828 0l3.536 3.536a2 2 0 0 1 0 2.828L10.614 21.928a2 2 0 0 1-2.828 0L4.25 18.392Zm2.121-2.121l-.707.707L9.2 20.513L20.514 9.2l-3.536-3.536l-.707.707l1.355 1.356a1 1 0 0 1-1.414 1.414l-1.355-1.355l-.707.707l.53.53a1 1 0 0 1-1.414 1.414l-.53-.53l-.708.707l1.355 1.355a1 1 0 1 1-1.414 1.414l-1.355-1.355l-.707.707l.53.53a1 1 0 1 1-1.414 1.415l-.53-.53l-.707.707l1.355 1.355a1 1 0 1 1-1.414 1.414L6.37 16.271Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}