package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10 12h3v-1h-3v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M9.5 0A1.5 1.5 0 0 0 8 1.5V3H1.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H6v2H4v1h9.5a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-4ZM8.085 14H7v-2h1v1.5c0 .175.03.344.085.5ZM9.5 14h4a.5.5 0 0 0 .5-.5V6H9v7.5a.5.5 0 0 0 .5.5ZM9 5h5V1.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}