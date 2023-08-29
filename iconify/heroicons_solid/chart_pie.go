package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M2 10a8 8 0 0 1 8-8v8h8a8 8 0 1 1-16 0Z"/><path d="M12 2.252A8.014 8.014 0 0 1 17.748 8H12V2.252Z"/></g>`),
		g.Group(children),
	)
}