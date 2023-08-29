package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorFillAccentTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.512 17.75a2.624 2.624 0 0 0-.012.25a3.5 3.5 0 0 0 6.973.438c.479.456.777 1.1.777 1.812v1.5a2.5 2.5 0 0 1-2.5 2.5H6.25a2.5 2.5 0 0 1-2.5-2.5v-1.5a2.5 2.5 0 0 1 2.5-2.5h1.011l1.102 1.102a3.75 3.75 0 0 0 5.303 0l1.101-1.102h1.745Z"/>`),
		g.Group(children),
	)
}