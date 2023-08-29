package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureLargeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 16q-.425 0-.713-.288T6 15V7q0-.425.288-.713T7 6h10q.425 0 .713.288T18 7v8q0 .425-.288.713T17 16H7Zm-4 4q-.425 0-.713-.288T2 19q0-.425.288-.713T3 18h17V5q0-.425.288-.713T21 4q.425 0 .713.288T22 5v13q0 .825-.588 1.413T20 20H3Z"/>`),
		g.Group(children),
	)
}