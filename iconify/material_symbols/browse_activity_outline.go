package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseActivityOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v4h-2V5H4v4H2Zm2 9q-.825 0-1.413-.588T2 16v-5h2v5h16v-5h2v5q0 .825-.588 1.413T20 18H4Zm-3 3v-2h22v2H1Zm11-10.5ZM2 11V9h6q.275 0 .525.15t.375.4l1.175 2.325L13.15 6.5q.125-.225.35-.363T14 6q.275 0 .525.138t.375.412L16.125 9H22v2h-6.5q-.275 0-.525-.138t-.375-.412l-.65-1.325l-3.075 5.375q-.125.25-.375.375T9.975 15q-.275 0-.513-.15t-.362-.4L7.375 11H2Z"/>`),
		g.Group(children),
	)
}