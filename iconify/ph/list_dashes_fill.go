package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDashesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 56v16a8 8 0 0 1-8 8H96a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h120a8 8 0 0 1 8 8ZM56 48H40a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h16a8 8 0 0 0 8-8V56a8 8 0 0 0-8-8Zm160 64H96a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h120a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm-160 0H40a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h16a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm160 64H96a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h120a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm-160 0H40a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h16a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}