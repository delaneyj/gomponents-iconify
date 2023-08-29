package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 14.5h7m-3.5 0v-5m0 0a4 4 0 0 0 4-4v-4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v4a4 4 0 0 0 4 4Zm-4-7h-1a2 2 0 1 0 0 4h1m8-4h1a2 2 0 1 1 0 4h-1"/>`),
		g.Group(children),
	)
}