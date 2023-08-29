package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 16h12l-4-8V1h1V0H5v1h1v7l-4 8zM9 1v7.2l1.9 3.8H5.1L7 8.2V1h2z"/>`),
		g.Group(children),
	)
}