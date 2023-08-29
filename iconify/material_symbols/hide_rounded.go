package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 16.4l-3.875 3.9q-.3.3-.7.3t-.7-.3q-.3-.3-.3-.713t.3-.712L7.6 15H6q-.425 0-.713-.287T5 14q0-.425.288-.713T6 13h4q.425 0 .713.288T11 14v4q0 .425-.288.713T10 19q-.425 0-.713-.288T9 18v-1.6ZM16.4 9H18q.425 0 .713.288T19 10q0 .425-.288.713T18 11h-4q-.425 0-.713-.288T13 10V6q0-.425.288-.713T14 5q.425 0 .713.288T15 6v1.6l3.875-3.9q.3-.3.7-.3t.7.3q.3.3.3.712t-.3.713L16.4 9Z"/>`),
		g.Group(children),
	)
}