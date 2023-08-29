package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentFireLitFlameTorchTrending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.15.53a.39.39 0 0 0-.4 0a.26.26 0 0 0-.06.34C6.92 3 7.18 5.9 5.5 7.5a5.52 5.52 0 0 1-1.5-2A3.89 3.89 0 0 0 2 9a4.7 4.7 0 0 0 5 4.5c3.22 0 4.89-2 5-4.5c.13-3-2-6.69-5.85-8.47Z"/><path d="M9.5 9a2 2 0 0 1-2 2"/></g>`),
		g.Group(children),
	)
}