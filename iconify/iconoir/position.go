package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Position(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0 0v2m-7-9H3m9-7V3m7 9h2"/>`),
		g.Group(children),
	)
}