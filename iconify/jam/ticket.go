package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 9H8a1 1 0 1 1 0-2h2V2.001a4.999 4.999 0 0 1-8 0V7h2a1 1 0 1 1 0 2H2v8.999a4.999 4.999 0 0 1 8 0V9zM0 20V0h3.17a3.001 3.001 0 0 0 5.66 0H12v20H8.83a3.001 3.001 0 0 0-5.66 0H0z"/>`),
		g.Group(children),
	)
}