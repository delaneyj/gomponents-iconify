package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.101 13.9a6.004 6.004 0 0 1-.601.08V13h-1v.98a6.001 6.001 0 0 1-5.482-5.518H3v-1h-.976A6.001 6.001 0 0 1 7.5 2.02V3h1v-.98a6.001 6.001 0 0 1 5.48 5.48H13v1h.98a6.004 6.004 0 0 1-.08.601c.334.077.652.196.95.351a7 7 0 1 0-5.397 5.397a3.973 3.973 0 0 1-.352-.95Zm.803-3.433L6.99 9.01L4.967 4.967L9.009 6.99l1.459 2.913a4.02 4.02 0 0 0-.564.563Zm-.469-1.032L8.481 7.52l-1.916-.955l.954 1.917l1.916.954Z" clip-rule="evenodd"/><circle cx="13" cy="13" r="3"/></g>`),
		g.Group(children),
	)
}