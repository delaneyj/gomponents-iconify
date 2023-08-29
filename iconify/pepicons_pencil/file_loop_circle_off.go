package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLoopCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M8.5 7a2 2 0 0 1 2-2h5.1a.5.5 0 0 1 .35.144l4.4 4.333a.5.5 0 0 1 .15.356V17a2 2 0 0 1-2 2h-9a.5.5 0 0 1 0-1h9a1 1 0 0 0 1-1v-6.667h-2.9a1.5 1.5 0 0 1-1.5-1.5V6h-4.6a1 1 0 0 0-1 1v2.5a.5.5 0 0 1-1 0V7Zm7.6-.306l2.68 2.64H16.6a.5.5 0 0 1-.5-.5v-2.14Z"/><path d="M10.998 11.628a2.291 2.291 0 1 0-2.15 4.047a2.291 2.291 0 0 0 2.15-4.047Zm-3.981.48a3.291 3.291 0 1 1 1.82 4.652l-1.61 3.03a.5.5 0 1 1-.883-.47l1.61-3.03a3.292 3.292 0 0 1-.937-4.183Z"/></g><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}