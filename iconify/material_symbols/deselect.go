package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deselect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19h2v2Zm2 0v-2h2v2H7Zm4 0v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 0V3q.825 0 1.413.588T21 5h-2Zm.775 17.6l-5.6-5.6H7V9.825l-5.6-5.6L2.8 2.8l18.4 18.4l-1.425 1.4Zm.05-5.6L19 16.175V15h2v2h-1.175ZM9 15h3.175L9 11.825V15Zm8-.825l-2-2V9h-3.175l-2-2H17v7.175ZM7.825 5L7 4.175V3h2v2H7.825ZM3 17v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm16 4v-2h2v2h-2Zm0-4V7h2v2h-2Z"/>`),
		g.Group(children),
	)
}