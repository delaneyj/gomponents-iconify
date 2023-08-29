package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridDotsTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 17a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm7 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4ZM5 17a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm7-7a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm7 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4ZM5 10a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm7-7a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm7 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4ZM5 3a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}