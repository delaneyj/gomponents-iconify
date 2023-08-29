package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterAltOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.8 11.975l-1.425-1.425L16.95 6H8.825l-2-2H19q.625 0 .9.55t-.1 1.05l-5 6.375Zm-.8 4.85V19q0 .425-.288.713T13 20h-2q-.425 0-.713-.288T10 19v-6.175l-7.9-7.9q-.275-.275-.275-.687t.275-.713q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288L14 16.825Zm-.625-6.275Z"/>`),
		g.Group(children),
	)
}