package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M24 21a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Zm0 2a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Z"/><path d="M24.384 9.91a1 1 0 0 1 .914 1.08l-.602 7.187a1 1 0 0 1-1.994-.167l.603-7.187a1 1 0 0 1 1.08-.913ZM14.5 40a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Zm0 2a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Z"/><path d="M17.757 36.757a1 1 0 0 1-1.414 0l-5.1-5.1a1 1 0 0 1 1.414-1.414l5.1 5.1a1 1 0 0 1 0 1.414ZM33.5 38a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Zm0 2a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Z"/><path d="M37.759 29.745a1 1 0 0 1-.544 1.306l-6.668 2.748a1 1 0 0 1-.762-1.85l6.668-2.748a1 1 0 0 1 1.306.544Z"/></g>`),
		g.Group(children),
	)
}