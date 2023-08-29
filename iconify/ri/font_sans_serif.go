package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSansSerif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4h14v4h-1.5C17 6 17 5 15 5h-5v7h3c1 0 2-.5 2-2h1v5h-1c0-1.5-1-2-2-2h-3v4.5c0 2.5 3.5 2.5 3.5 2.5v1H5v-1c2-.5 2-1.5 2-2.5v-10c0-1 0-2-2-2.5V4Z"/>`),
		g.Group(children),
	)
}