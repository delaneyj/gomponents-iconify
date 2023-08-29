package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.83 117.17l-96-96a4 4 0 0 0-5.66 0l-96 96A4 4 0 0 0 32 124h44v28a4 4 0 0 0 4 4h96a4 4 0 0 0 4-4v-28h44a4 4 0 0 0 2.83-6.83ZM176 116a4 4 0 0 0-4 4v28H84v-28a4 4 0 0 0-4-4H41.66L128 29.66L214.34 116Zm4 100a4 4 0 0 1-4 4H80a4 4 0 0 1 0-8h96a4 4 0 0 1 4 4Zm0-32a4 4 0 0 1-4 4H80a4 4 0 0 1 0-8h96a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}