package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 10.1C16 5.7 12.4 2 8 2s-8 3.7-8 8.1c0 1.4.3 2.9.9 3.9h4.9c.5.6 1.3 1 2.2 1s1.7-.4 2.2-1h4.9c.6-1 .9-2.5.9-3.9zM14 7v1l-4.1 3.5c0 .1.1.3.1.5c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2c.3 0 .6.1.8.2L13 7h1zm-4-3h1v1h-1V4zM5 4h1v1H5V4zm-3 8H1v-1h1v1zm1-4H2V7h1v1zm12 4h-1v-1h1v1z"/><path fill="currentColor" d="M9 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}