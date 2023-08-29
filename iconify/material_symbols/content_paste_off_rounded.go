package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentPasteOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18.15l-2-2V5h-2v2q0 .425-.288.713T16 8h-5.15l-5-5h3.325q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H19q.825 0 1.413.588T21 5v13.15ZM12 5q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Zm4.15 14L5 7.85V19h11.15ZM5 21q-.825 0-1.413-.588T3 19V5.85l-.925-.925q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3l-.925-.9H5Z"/>`),
		g.Group(children),
	)
}