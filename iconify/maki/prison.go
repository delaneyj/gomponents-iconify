package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prison(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.5 1v13H12V1H3.5zm6 1H11v3.5H9.5V2zm-5 .055H6V7H4.5V2.055zm2.5 0h1.5V7H7V2.055zM10.25 6.5a.75.75 0 0 1 0 1.5a.75.75 0 0 1 0-1.5zM7 8h1.473l.027 5H7.027L7 8zm-2.5.166H6V13H4.5V8.166zM9.5 9H11v4H9.5V9z"/>`),
		g.Group(children),
	)
}