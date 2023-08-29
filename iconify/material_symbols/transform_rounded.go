package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransformRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9H3q-.425 0-.713-.288T2 8q0-.425.288-.713T3 7h4V4.85l-.875.875q-.3.3-.713.288T4.7 5.7q-.275-.3-.288-.7t.288-.7l2.6-2.6q.3-.3.7-.3t.7.3l2.6 2.6q.3.3.287.7t-.287.7q-.3.3-.713.313t-.712-.288L9 4.85V15h12q.425 0 .713.288T22 16q0 .425-.288.713T21 17h-4v2.15l.875-.875q.3-.3.713-.288t.712.313q.275.3.288.7t-.288.7l-2.6 2.6q-.3.3-.7.3t-.7-.3l-2.6-2.6q-.3-.3-.287-.7t.287-.7q.3-.3.713-.312t.712.287l.875.875V17H9q-.825 0-1.413-.588T7 15V9Zm8 4V9h-4V7h4q.825 0 1.413.588T17 9v4h-2Z"/>`),
		g.Group(children),
	)
}