package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitRequestCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M12.25 7.25v3m-8.5-4.5v4.5m10.5-8.5l-3.5 3.5m0-3.5l3.5 3.5"/></g>`),
		g.Group(children),
	)
}