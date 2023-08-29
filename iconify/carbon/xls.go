package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 23h-6v-2h6v-4h-4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h6v2h-6v4h4a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm-14-2V9h-2v14h8v-2h-6zM10 9H8l-2 6l-2-6H2l2.752 7L2 23h2l2-6l2 6h2l-2.755-7L10 9z"/>`),
		g.Group(children),
	)
}