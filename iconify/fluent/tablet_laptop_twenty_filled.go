package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletLaptopTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 3a2 2 0 0 0-2 2v3h6a3 3 0 0 1 3 3v1h1a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H7Zm10.5 11H14v-1h3.5a.5.5 0 0 1 0 1Zm-11 0a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2ZM2 11a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4Z"/>`),
		g.Group(children),
	)
}