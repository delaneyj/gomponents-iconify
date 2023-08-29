package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.707 0H2.5A1.5 1.5 0 0 0 1 1.5v12A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4.293L9.707 0ZM7 6H4V5h4v7H7V6ZM6 9H4V8h2v1Zm-2 3h2v-1H4v1Zm7-6H9V5h2v1ZM9 9h2V8H9v1Zm2 3H9v-1h2v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}