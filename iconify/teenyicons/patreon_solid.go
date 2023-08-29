package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 0H0v15h3V0Zm6.5 0a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11Z"/>`),
		g.Group(children),
	)
}