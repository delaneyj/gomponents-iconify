package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 1.5V16h3v2h-3v4h-2v-4H4v-1.102L14 1.5h2ZM14 16V5.171L6.968 16H14Z"/>`),
		g.Group(children),
	)
}