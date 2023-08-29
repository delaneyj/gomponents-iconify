package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AidOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 9v10V9Zm5 5H7h10ZM1 6.995C1 5.893 1.89 5 2.991 5H21.01C22.109 5 23 5.893 23 6.995v14.01C23 22.107 22.11 23 21.009 23H2.99A1.992 1.992 0 0 1 1 21.005V6.995ZM7 5V2.01C7 1.451 7.456 1 7.995 1h8.01c.55 0 .995.443.995 1.01V5"/>`),
		g.Group(children),
	)
}