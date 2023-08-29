package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.5 4a2 2 0 0 1 2-2h5.1a.5.5 0 0 1 .35.144l4.4 4.333a.5.5 0 0 1 .15.356V14a2 2 0 0 1-2 2h-9a.5.5 0 0 1 0-1h9a1 1 0 0 0 1-1V7.333h-2.9a1.5 1.5 0 0 1-1.5-1.5V3H7.5a1 1 0 0 0-1 1v2.5a.5.5 0 0 1-1 0V4Zm7.6-.306l2.68 2.64H13.6a.5.5 0 0 1-.5-.5v-2.14Z"/><path d="M7.998 8.628a2.291 2.291 0 1 0-2.15 4.047a2.291 2.291 0 0 0 2.15-4.047Zm-3.981.48a3.291 3.291 0 1 1 1.82 4.652l-1.61 3.03a.5.5 0 1 1-.883-.47l1.61-3.03a3.292 3.292 0 0 1-.937-4.183Z"/></g>`),
		g.Group(children),
	)
}