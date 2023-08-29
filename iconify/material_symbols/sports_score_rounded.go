package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsScoreRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6h2V4H9v2Zm4 0V4h2v2h-2Zm-4 8v-2h2v2H9Zm8-4V8h2v2h-2Zm0 4v-2h2v2h-2Zm-4 0v-2h2v2h-2Zm4-8V4h2v2h-2Zm-6 2V6h2v2h-2ZM6 20q-.425 0-.713-.288T5 19V5q0-.425.288-.713T6 4q.425 0 .713.288T7 5v1h2v2H7v2h2v2H7v7q0 .425-.288.713T6 20Zm9-8v-2h2v2h-2Zm-4 0v-2h2v2h-2Zm-2-2V8h2v2H9Zm4 0V8h2v2h-2Zm2-2V6h2v2h-2Z"/>`),
		g.Group(children),
	)
}