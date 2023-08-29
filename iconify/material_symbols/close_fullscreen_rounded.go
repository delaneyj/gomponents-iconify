package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseFullscreenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 15.4l-5.9 5.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7L8.6 14H5q-.425 0-.713-.288T4 13q0-.425.288-.713T5 12h6q.425 0 .713.288T12 13v6q0 .425-.288.713T11 20q-.425 0-.713-.288T10 19v-3.6Zm5.4-5.4H19q.425 0 .713.288T20 11q0 .425-.288.713T19 12h-6q-.425 0-.713-.288T12 11V5q0-.425.288-.713T13 4q.425 0 .713.288T14 5v3.6l5.9-5.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L15.4 10Z"/>`),
		g.Group(children),
	)
}