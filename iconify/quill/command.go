package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Command(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 19v4a4 4 0 1 0 4-4h-4Zm0 0v-6m0 6h-6m6-6V9a4 4 0 1 1 4 4h-4Zm0 0h-6m-4 0h-.007m0 0A4 4 0 1 1 13 9v4m-4.007 0H13m0 0v6m0 0v4a4 4 0 1 1-4-4h4Z"/>`),
		g.Group(children),
	)
}