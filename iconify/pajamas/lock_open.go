package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.523.029a4 4 0 0 1 4.37 3.05a.75.75 0 0 1-1.46.345a2.5 2.5 0 0 0-4.933.592V7H13a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1.75H4V4.024A4 4 0 0 1 7.523.029ZM3.5 8.5v6h9v-6h-9ZM9 10H7v3h2v-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}