package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagnosisOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19h6q.425 0 .713-.288T16 18q0-.425-.288-.713T15 17H9q-.425 0-.713.288T8 18q0 .425.288.713T9 19Zm0-3h6q.425 0 .713-.288T16 15q0-.425-.288-.713T15 14H9q-.425 0-.713.288T8 15q0 .425.288.713T9 16Zm3-10.2q-.3-.375-.788-.588T10.2 5q-.9 0-1.55.65T8 7.2q0 1.325 1.225 2.513t2.1 1.987q.275.275.675.275t.675-.275q.875-.8 2.1-1.987T16 7.2q0-.9-.65-1.55T13.8 5q-.525 0-1.013.213T12 5.8ZM18 22H6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22ZM6 20h12V4H6v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}