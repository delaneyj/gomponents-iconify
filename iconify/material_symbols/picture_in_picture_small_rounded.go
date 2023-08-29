package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureSmallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.425 0-.713-.288T2 19q0-.425.288-.713T3 18h17V5q0-.425.288-.713T21 4q.425 0 .713.288T22 5v13q0 .825-.588 1.413T20 20H3Zm8-4q-.425 0-.713-.288T10 15v-4q0-.425.288-.713T11 10h6q.425 0 .713.288T18 11v4q0 .425-.288.713T17 16h-6Z"/>`),
		g.Group(children),
	)
}