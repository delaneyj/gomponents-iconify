package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5.5h-2v14h2V.5Zm2 5a5 5 0 1 0 10 0a5 5 0 0 0-10 0Z"/>`),
		g.Group(children),
	)
}