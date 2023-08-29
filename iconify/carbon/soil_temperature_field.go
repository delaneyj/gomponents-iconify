package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoilTemperatureField(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 16v-5h1a4.005 4.005 0 0 0 4-4V4h-3a3.978 3.978 0 0 0-2.747 1.107A6.003 6.003 0 0 0 5 2H2v3a6.007 6.007 0 0 0 6 6h1v5H2v2h14v-2zm2-10h1v1a2.002 2.002 0 0 1-2 2h-1V8a2.002 2.002 0 0 1 2-2zM8 9a4.005 4.005 0 0 1-4-4V4h1a4.005 4.005 0 0 1 4 4v1zM2 21h14v2H2zm0 5h14v2H2zm23 4a4.986 4.986 0 0 1-3-8.98V15a3 3 0 0 1 6 0v6.02A4.986 4.986 0 0 1 25 30zm0-16a1.001 1.001 0 0 0-1 1v7.13l-.497.29A2.968 2.968 0 0 0 22 25a3 3 0 0 0 6 0a2.968 2.968 0 0 0-1.503-2.581L26 22.13V15a1.001 1.001 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}