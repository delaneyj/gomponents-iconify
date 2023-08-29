package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11q-.425 0-.713-.288T2 10V5q0-.425.288-.713T3 4h18q.425 0 .713.288T22 5v5q0 .425-.288.713T21 11H3Zm8 9q-.425 0-.713-.288T10 19v-5q0-.425.288-.713T11 13h10q.425 0 .713.288T22 14v5q0 .425-.288.713T21 20H11Zm-8 0q-.425 0-.713-.288T2 19v-5q0-.425.288-.713T3 13h4q.425 0 .713.288T8 14v5q0 .425-.288.713T7 20H3Z"/>`),
		g.Group(children),
	)
}