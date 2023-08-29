package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.101 13.9a6.004 6.004 0 0 1-.601.08V13h-1v.98a6.001 6.001 0 0 1-5.482-5.518H3v-1h-.976A6.001 6.001 0 0 1 7.5 2.02V3h1v-.98a6.001 6.001 0 0 1 5.48 5.48H13v1h.98a6.004 6.004 0 0 1-.08.601c.334.077.652.196.95.351a7 7 0 1 0-5.397 5.397a3.973 3.973 0 0 1-.352-.95Zm.803-3.433L6.99 9.01L4.967 4.967L9.009 6.99l1.459 2.913a4.02 4.02 0 0 0-.564.563Zm-.469-1.032L8.481 7.52l-1.916-.955l.954 1.917l1.916.954Z"/><path d="M11.333 10.506a3 3 0 1 1 3.333 4.987a3 3 0 0 1-3.333-4.987Zm1.698 3.817l1.79-2.387l-.8-.6l-1.48 1.974l-.876-.7l-.624.78l1.278 1.023l.712-.09Z"/></g>`),
		g.Group(children),
	)
}