package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.075 21.9l-.925-.9H5q-.825 0-1.413-.588T3 19V5.85l-.925-.925q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM5 19h11.15l-4.875-4.875L10.4 15q-.275.275-.7.275T9 15q-.275-.275-.275-.7T9 13.6l.875-.875L5 7.85V19ZM7.85 5l-2-2H11q.425 0 .713.288T12 4q0 .425-.288.713T11 5H7.85Zm6.275 6.275l-1.4-1.4L17.6 5H15q-.425 0-.713-.288T14 4q0-.425.288-.713T15 3h5q.425 0 .713.288T21 4v5q0 .425-.288.713T20 10q-.425 0-.713-.288T19 9V6.4l-4.875 4.875ZM21 18.15l-2-2V13q0-.425.288-.713T20 12q.425 0 .713.288T21 13v5.15Z"/>`),
		g.Group(children),
	)
}