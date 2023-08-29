package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 13.565C2 11.512 4 11 6 11v9a4 4 0 0 1-4-4v-2.435Zm20 0C22 11.512 20 11 18 11v9a4 4 0 0 0 4-4v-2.435ZM6 20V10a6 6 0 1 1 12 0v10"/>`),
		g.Group(children),
	)
}