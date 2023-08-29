package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.25 3.75a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM10.45 5a2.75 2.75 0 1 0-4.9 0H3.82a1 1 0 0 0-.98.804l-1.6 8A1 1 0 0 0 2.22 15h11.56a1 1 0 0 0 .98-1.196l-1.6-8A1 1 0 0 0 12.18 5h-1.73ZM8 6.5H4.23l-1.4 7h10.34l-1.4-7H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}