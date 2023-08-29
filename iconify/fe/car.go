package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCar1" fill="currentColor" fill-rule="nonzero"><path id="feCar2" d="M3 18v-5a2 2 0 0 1 2-2V8a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v3a2 2 0 0 1 2 2v5a2 2 0 1 1-4 0H7a2 2 0 1 1-4 0ZM9 6a2 2 0 0 0-2 2v3h10V8a2 2 0 0 0-2-2H9Zm-3 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm12 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}