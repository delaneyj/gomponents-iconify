package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBuilding0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBuilding1" fill="currentColor" fill-rule="nonzero"><path id="feBuilding2" d="M6 2h12a2 2 0 0 1 2 2v18H4V4a2 2 0 0 1 2-2Zm0 18h4v-4h4v4h4V4H6v16Zm7-14h3v3h-3V6Zm-5 5h3v3H8v-3Zm5 0h3v3h-3v-3ZM8 6h3v3H8V6Z"/></g></g>`),
		g.Group(children),
	)
}