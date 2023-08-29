package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandRadixUi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0M6 3h5v5H6zm5 8v10a5 5 0 0 1-.217-9.995L11 11z"/>`),
		g.Group(children),
	)
}