package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopRulerCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M4.25 18.392a2 2 0 0 1 0-2.828L15.564 4.25a2 2 0 0 1 2.828 0l3.536 3.536a2 2 0 0 1 0 2.828L10.614 21.928a2 2 0 0 1-2.828 0L4.25 18.392Zm2.121-2.121l-.707.707L9.2 20.513L20.514 9.2l-3.536-3.536l-.707.707l1.355 1.356a1 1 0 0 1-1.414 1.414l-1.355-1.355l-.707.707l.53.53a1 1 0 0 1-1.414 1.414l-.53-.53l-.708.707l1.355 1.355a1 1 0 1 1-1.414 1.414l-1.355-1.355l-.707.707l.53.53a1 1 0 1 1-1.414 1.415l-.53-.53l-.707.707l1.355 1.355a1 1 0 1 1-1.414 1.414L6.37 16.271Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopRulerCircleFilled0)"/></g>`),
		g.Group(children),
	)
}