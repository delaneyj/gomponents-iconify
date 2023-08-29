package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 9H8a2.002 2.002 0 0 0-2 2v10a2.002 2.002 0 0 0 2 2h13a2.002 2.002 0 0 0 2-2V11a2.002 2.002 0 0 0-2-2Zm0 4h-3v-2h3Zm-3 2h3v2h-3Zm-2-4v10h-3v-7a1 1 0 0 0-1-1H8v-2Zm-5 6H8v-2h3Zm-3 2h3v2H8Zm10 2v-2h3v2Z"/><path fill="currentColor" d="M23.532 27H4a2.002 2.002 0 0 1-2-2V7a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v12.638a2.004 2.004 0 0 1-.464 1.28l-4.468 5.362a1.994 1.994 0 0 1-1.536.72ZM4 7v18h19.532L28 19.638V7Z"/>`),
		g.Group(children),
	)
}