package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5 6v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3Zm3 15a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H8Z" clip-rule="evenodd"/><path d="M8 10.5v-4a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5Zm7 8.2v-4.9a.8.8 0 0 1 .8-.8h1.4a.8.8 0 0 1 .8.8v4.9a.8.8 0 0 1-.8.8h-1.4a.8.8 0 0 1-.8-.8ZM8 15v-1.5a.5.5 0 0 1 .5-.5H10a.5.5 0 0 1 .5.5V15a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 15Zm3.5 0v-1.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5V15a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5ZM8 19v-1.5a.5.5 0 0 1 .5-.5H10a.5.5 0 0 1 .5.5V19a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 19Zm3.5 0v-1.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5V19a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5Z"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}