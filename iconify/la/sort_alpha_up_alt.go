package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlphaUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v2h5.56l-5.28 5.28l-.28.31V15h8v-2H7.44l5.28-5.28l.28-.31V5H5zm18 .5l-.72.69L18 10.5l1.41 1.41L22 9.31V28h2V9.31l2.59 2.6L28 10.5l-4.28-4.31L23 5.5zM8.19 17l-.22.66L6.03 23H6v.06l-.94 2.6l-.06.15V27h2v-.84L7.41 25h3.18l.41 1.16V27h2v-1.19l-.06-.15l-.94-2.6V23h-.03l-1.94-5.34l-.22-.66H8.19zM9 20.66L9.84 23H8.16L9 20.66z"/>`),
		g.Group(children),
	)
}