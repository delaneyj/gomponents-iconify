package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feWineGlass0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWineGlass1" fill="currentColor"><path id="feWineGlass2" d="M9 20h2v-6a4 4 0 0 1-4-4V4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v6a4 4 0 0 1-4 4v6h2a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2ZM9 4v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4H9Z"/></g></g>`),
		g.Group(children),
	)
}