package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 5a4 4 0 1 1 8 0v1h1a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1V5Zm6.5 0v1h-5V5a2.5 2.5 0 0 1 5 0Zm-7 2.5v6h9v-6h-9ZM9 12V9H7v3h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}