package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelMini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 14.172V10a2 2 0 0 0-2-2h-4.172a2 2 0 0 0-1.414.586l-8 8a2 2 0 0 0 0 2.828l4.172 4.172a2 2 0 0 0 2.828 0l8-8A2 2 0 0 0 24 14.172Z"/>`),
		g.Group(children),
	)
}