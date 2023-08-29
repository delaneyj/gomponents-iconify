package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReflectHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.386 15.21l9-7A1 1 0 0 1 30 9v14a1 1 0 0 1-1.614.79l-9-7a1 1 0 0 1 0-1.58ZM17 30h-2V2h2zm-4-14a1.001 1.001 0 0 1-.386.79l-9 7A1 1 0 0 1 2 23V9a1 1 0 0 1 1.614-.79l9 7A1.001 1.001 0 0 1 13 16Zm-9 4.956L10.371 16L4 11.044Z"/>`),
		g.Group(children),
	)
}