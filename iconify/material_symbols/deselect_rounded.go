package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeselectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19h2v2Zm2 0v-2h2v2H7Zm4 0v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 0V3q.825 0 1.413.588T21 5h-2Zm.075 16.9l-4.9-4.9H8q-.425 0-.713-.288T7 16V9.825l-4.9-4.9q-.275-.275-.275-.687t.275-.713q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288ZM9 15h3.175L9 11.825V15Zm8-.825l-2-2V9h-3.175l-2-2H16q.425 0 .713.288T17 8v6.175ZM7.825 5L7 4.175V3h2v2H7.825Zm12 12L19 16.175V15h2v2h-1.175ZM3 17v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm16 4v-2h2v2h-2Zm0-4V7h2v2h-2Z"/>`),
		g.Group(children),
	)
}