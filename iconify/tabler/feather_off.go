package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeatherOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m4 20l8-8m2-7v5h5M9 11v4h4m-7-2v5h5m-5-5l3.502-3.502m2.023-2.023L14 5"/><path d="M19 10c.638-.636 1-1.515 1-2.486A3.515 3.515 0 0 0 16.483 4c-.97 0-1.847.367-2.483 1m-3 13l3.499-3.499m2.008-2.008L19 10M3 3l18 18"/></g>`),
		g.Group(children),
	)
}