package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballField(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 5h9.4a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H12m0-14H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6H12m0-14v14"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6ZM2 17A5 5 0 0 0 2 7m20 10a5 5 0 0 1 0-10"/></g>`),
		g.Group(children),
	)
}