package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SavingsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11q.425 0 .713-.288T17 10q0-.425-.288-.713T16 9q-.425 0-.713.288T15 10q0 .425.288.713T16 11ZM9 9h3q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7H9q-.425 0-.713.288T8 8q0 .425.288.713T9 9ZM5.925 21q-.575 0-1.112-.4t-.713-.975q-.625-2.1-1.025-3.638t-.638-2.7Q2.2 12.126 2.1 11.226T2 9.5q0-2.3 1.6-3.9T7.5 4h5q.675-.9 1.713-1.45T16.5 2q.625 0 1.063.438T18 3.5q0 .15-.038.3t-.087.275q-.1.275-.188.55t-.137.6L19.825 7.5H21q.425 0 .713.288T22 8.5v5.25q0 .325-.188.575t-.512.375l-2.125.7l-1.25 4.175q-.2.65-.725 1.038T16 21h-2q-.825 0-1.413-.588T12 19h-2q0 .825-.588 1.413T8 21H5.925Z"/>`),
		g.Group(children),
	)
}