package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightShelterOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Zm0-2h12v-9l-6-4.5L6 10v9Zm6-6.75ZM8 16.5h8v1q0 .2.15.35t.35.15q.2 0 .35-.15t.15-.35V14q0-.825-.588-1.413T15 12h-2.5q-.425 0-.713.288T11.5 13v2.5H8v-4q0-.2-.15-.35T7.5 11q-.2 0-.35.15T7 11.5v6q0 .2.15.35t.35.15q.2 0 .35-.15T8 17.5v-1ZM9.75 15q.525 0 .888-.363T11 13.75q0-.525-.363-.888T9.75 12.5q-.525 0-.888.363t-.362.887q0 .525.363.888T9.75 15Z"/>`),
		g.Group(children),
	)
}