package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 140a4 4 0 0 0-4 4a44.05 44.05 0 0 1-44 44a44 44 0 0 0-36 18.73V124h36a4 4 0 0 0 0-8h-36V83.71a28 28 0 1 0-8 0V116H88a4 4 0 0 0 0 8h36v82.73A44 44 0 0 0 88 188a44.05 44.05 0 0 1-44-44a4 4 0 0 0-8 0a52.06 52.06 0 0 0 52 52a36 36 0 0 1 36 36a4 4 0 0 0 8 0a36 36 0 0 1 36-36a52.06 52.06 0 0 0 52-52a4 4 0 0 0-4-4ZM108 56a20 20 0 1 1 20 20a20 20 0 0 1-20-20Z"/>`),
		g.Group(children),
	)
}