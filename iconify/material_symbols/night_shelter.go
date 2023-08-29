package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightShelter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18h1v-1.5h8V18h1v-4q0-.825-.588-1.413T15 12h-3.5v3.5H8V11H7v7Zm2.75-3q.525 0 .888-.363T11 13.75q0-.525-.363-.888T9.75 12.5q-.525 0-.888.363t-.362.887q0 .525.363.888T9.75 15ZM4 21V9l8-6l8 6v12H4Z"/>`),
		g.Group(children),
	)
}