package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResetImageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9H4q-.425 0-.713-.288T3 8V4q0-.425.288-.713T4 3q.425 0 .713.288T5 4v2.35Q6.25 4.8 8.063 3.9T12 3q2.65 0 4.763 1.338T20 7.875q.2.4-.075.763T19.15 9q-.35 0-.65-.163t-.5-.462q-.925-1.525-2.5-2.45T12 5q-1.425 0-2.688.525T7.1 7H8q.425 0 .713.288T9 8q0 .425-.288.713T8 9Zm-1 9h10q.3 0 .45-.275t-.05-.525l-2.75-3.675q-.15-.2-.4-.2t-.4.2L11.25 17L9.4 14.525q-.15-.2-.4-.2t-.4.2l-2 2.675q-.2.25-.05.525T7 18Zm-2 4q-.825 0-1.413-.588T3 20v-7q0-.425.288-.713T4 12q.425 0 .713.288T5 13v7h14v-7q0-.425.288-.713T20 12q.425 0 .713.288T21 13v7q0 .825-.588 1.413T19 22H5Z"/>`),
		g.Group(children),
	)
}