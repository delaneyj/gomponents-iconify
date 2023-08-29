package fontelico

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M425.518 1000c-567.358 0-567.356-1000 0-1000s567.358 1000 0 1000zm0-91.74c237.715 0 237.714-816.52 0-816.52s-237.715 816.52 0 816.52z"/>`),
		g.Group(children),
	)
}