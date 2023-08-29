package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyframeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 4a2.599 2.599 0 0 0-2 .957l-4.355 5.24a2.847 2.847 0 0 0-.007 3.598l4.368 5.256c.499.6 1.225.949 1.994.949a2.599 2.599 0 0 0 2-.957l4.355-5.24a2.847 2.847 0 0 0 .007-3.598l-4.368-5.256A2.593 2.593 0 0 0 12 4z"/></g>`),
		g.Group(children),
	)
}