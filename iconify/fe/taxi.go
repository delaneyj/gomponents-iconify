package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTaxi0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTaxi1" fill="currentColor" fill-rule="nonzero"><path id="feTaxi2" d="M8.314 5.059L9 3h6l.686 2.059A4.001 4.001 0 0 1 19 9v3a2 2 0 0 1 2 2v5a2 2 0 1 1-4 0H7a2 2 0 1 1-4 0v-5a2 2 0 0 1 2-2V9a4.001 4.001 0 0 1 3.314-3.941ZM9 7a2 2 0 0 0-2 2v3h10V9a2 2 0 0 0-2-2H9Zm-3 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm12 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}