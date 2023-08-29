package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCalculatorCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M5 6v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3Zm3 15a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H8Z" clip-rule="evenodd"/><path d="M8 10.5v-4a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5Zm7 8.2v-4.9a.8.8 0 0 1 .8-.8h1.4a.8.8 0 0 1 .8.8v4.9a.8.8 0 0 1-.8.8h-1.4a.8.8 0 0 1-.8-.8ZM8 15v-1.5a.5.5 0 0 1 .5-.5H10a.5.5 0 0 1 .5.5V15a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 15Zm3.5 0v-1.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5V15a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5ZM8 19v-1.5a.5.5 0 0 1 .5-.5H10a.5.5 0 0 1 .5.5V19a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 19Zm3.5 0v-1.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5V19a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCalculatorCircleFilled0)"/></g>`),
		g.Group(children),
	)
}