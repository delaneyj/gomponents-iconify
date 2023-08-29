package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.961 16.309a.5.5 0 0 0-.108-.163l-4.5-4.5a.5.5 0 0 0-.707.708L17.293 16H8.5A2.502 2.502 0 0 1 6 13.5v-10a.5.5 0 0 0-1 0v10A3.504 3.504 0 0 0 8.5 17h8.793l-3.647 3.646a.5.5 0 1 0 .708.708l4.5-4.5a.499.499 0 0 0 .146-.352V16.5a.5.5 0 0 0-.039-.191z"/>`),
		g.Group(children),
	)
}