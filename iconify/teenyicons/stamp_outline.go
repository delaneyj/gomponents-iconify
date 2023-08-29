package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StampOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 14.5h15m-8.5-8v3m2-3v3m-1-3a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-7 6v-1a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1H.5Z"/>`),
		g.Group(children),
	)
}