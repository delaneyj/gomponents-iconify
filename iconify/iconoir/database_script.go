package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseScript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 14V6a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v7"/><path d="M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/></g>`),
		g.Group(children),
	)
}