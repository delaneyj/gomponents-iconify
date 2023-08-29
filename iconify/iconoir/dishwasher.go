package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dishwasher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-2.5M21 7H3m9 9v5m0 0h-2m2 0h2"/><path d="M12 16c1.657 0 3-1.492 3-3.333V10H9v2.667C9 14.507 10.343 16 12 16Zm6-10.99l.01-.011M15 5.01l.01-.011"/></g>`),
		g.Group(children),
	)
}