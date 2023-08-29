package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18H8v-2h12v2ZM6 18H4v-2h2v2Zm14-5H8v-2h12v2ZM6 13H4v-2h2v2Zm14-5H8.023V6H20v2ZM6 8H4V6h2v2Z"/>`),
		g.Group(children),
	)
}