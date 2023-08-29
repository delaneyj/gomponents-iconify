package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 15V9H13v6h-1.5ZM6 15q-.45 0-.725-.313T5 14v-4q0-.375.275-.688T6 9h3q.45 0 .725.313T10 10v.5H6.5v3h2V12H10v2q0 .375-.275.688T9 15H6Zm8.5 0V9H19v1.5h-3v1h2V13h-2v2h-1.5Z"/>`),
		g.Group(children),
	)
}