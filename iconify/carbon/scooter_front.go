package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScooterFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 19h2v9h-2z"/><path fill="currentColor" d="M23 10V8h-3a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2H9v2h3v4.184A2.996 2.996 0 0 0 10 17v7h2v-7a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v7h2v-7a2.996 2.996 0 0 0-2-2.816V10Zm-5-2v6h-4V8Z"/>`),
		g.Group(children),
	)
}