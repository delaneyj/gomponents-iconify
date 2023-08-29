package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAdSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 6a1 1 0 0 0-1 1v1h2V7a1 1 0 0 0-1-1Zm6 2H9.5a.5.5 0 0 0 0 1H11V8Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 4.5A2.5 2.5 0 0 1 2.5 2h10A2.5 2.5 0 0 1 15 4.5v6a2.5 2.5 0 0 1-2.5 2.5h-10A2.5 2.5 0 0 1 0 10.5v-6ZM4 10V9h2v1h1V7a2 2 0 1 0-4 0v3h1Zm7-3H9.5a1.5 1.5 0 1 0 0 3H12V5h-1v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}