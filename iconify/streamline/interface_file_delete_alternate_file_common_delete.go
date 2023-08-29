package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFileDeleteAlternateFileCommonDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 7V1.5a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1H8"/><path d="M8.5.5v5h5M4.74 9.26L.5 13.5m0-4.24l4.24 4.24"/></g>`),
		g.Group(children),
	)
}