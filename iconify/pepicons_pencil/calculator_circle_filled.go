package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCalculatorCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M5.5 6v14A2.5 2.5 0 0 0 8 22.5h10a2.5 2.5 0 0 0 2.5-2.5V6A2.5 2.5 0 0 0 18 3.5H8A2.5 2.5 0 0 0 5.5 6ZM8 21.5A1.5 1.5 0 0 1 6.5 20V6A1.5 1.5 0 0 1 8 4.5h10A1.5 1.5 0 0 1 19.5 6v14a1.5 1.5 0 0 1-1.5 1.5H8Z" clip-rule="evenodd"/><path d="M8 10.5v-4a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5Zm7 8.2v-4.9a.8.8 0 0 1 .8-.8h1.4a.8.8 0 0 1 .8.8v4.9a.8.8 0 0 1-.8.8h-1.4a.8.8 0 0 1-.8-.8ZM8 15v-1.5a.5.5 0 0 1 .5-.5H10a.5.5 0 0 1 .5.5V15a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 15Zm3.5 0v-1.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5V15a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5ZM8 19v-1.5a.5.5 0 0 1 .5-.5H10a.5.5 0 0 1 .5.5V19a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 19Zm3.5 0v-1.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5V19a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCalculatorCircleFilled0)"/></g>`),
		g.Group(children),
	)
}