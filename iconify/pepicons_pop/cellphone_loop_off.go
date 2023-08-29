package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneLoopOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M12.675 15.137a.675.675 0 1 1-1.35 0a.675.675 0 0 1 1.35 0Z"/><path fill-rule="evenodd" d="M12 14.963a.175.175 0 1 0 0 .35a.175.175 0 0 0 0-.35Zm-1.175.175a1.175 1.175 0 1 1 2.35 0a1.175 1.175 0 0 1-2.35 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.5 3.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3H8A2.5 2.5 0 0 1 5.5 17v-.5a1 1 0 1 1 2 0v.5a.5.5 0 0 0 .5.5h7.5a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1h-7a1 1 0 0 0-1 1v2.25a1 1 0 0 1-2 0V3.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.347 9.38a2.193 2.193 0 1 0-2.567 3.555A2.193 2.193 0 0 0 7.347 9.38Zm-4.683-.677a4.193 4.193 0 1 1 1.827 6.342L2.4 17.942a1 1 0 1 1-1.622-1.17l2.091-2.897a4.195 4.195 0 0 1-.205-5.172Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}