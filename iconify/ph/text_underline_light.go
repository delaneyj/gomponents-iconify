package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 224a6 6 0 0 1-6 6H64a6 6 0 0 1 0-12h128a6 6 0 0 1 6 6Zm-70-26a62.07 62.07 0 0 0 62-62V56a6 6 0 0 0-12 0v80a50 50 0 0 1-100 0V56a6 6 0 0 0-12 0v80a62.07 62.07 0 0 0 62 62Z"/>`),
		g.Group(children),
	)
}