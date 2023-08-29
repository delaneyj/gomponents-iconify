package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTruck0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTruck1" fill="currentColor"><path id="feTruck2" d="M10 18a3 3 0 0 1-6 0a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11.001C16.105 4 17 4.895 17 5.999L20 6c2 0 2 2.896 2 4v6a2 2 0 0 1-2 2h-1a3 3 0 0 1-6 0h-3ZM4 6v9h11V6H4Zm13 2v4h3V8h-3Zm-1 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-9 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}