package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.83 117.17l-96-96a4 4 0 0 0-5.66 0l-96 96A4 4 0 0 0 32 124h44v84a12 12 0 0 0 12 12h80a12 12 0 0 0 12-12v-84h44a4 4 0 0 0 2.83-6.83ZM176 116a4 4 0 0 0-4 4v88a4 4 0 0 1-4 4H88a4 4 0 0 1-4-4v-88a4 4 0 0 0-4-4H41.66L128 29.66L214.34 116Z"/>`),
		g.Group(children),
	)
}