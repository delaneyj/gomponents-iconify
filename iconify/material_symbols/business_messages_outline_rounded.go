package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessMessagesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20q-.825 0-1.413-.588T5 18V8.75L1.7 4.825q-.2-.25-.075-.537T2.075 4H20q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H7ZM7 6v12h13V6H7Zm6.8 9H17q.425 0 .713-.288T18 14q0-.425-.288-.713T17 13h-5.575q-.675 0-.938.613T10.7 14.7l2.1 2.1q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-.4-.4Zm-.6-6H10q-.425 0-.713.288T9 10q0 .425.288.713T10 11h5.575q.675 0 .938-.613T16.3 9.3l-2.1-2.1q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7l.4.4ZM7 6v12V6Z"/>`),
		g.Group(children),
	)
}