package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftLockOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17v-4H5.1q-.65 0-.912-.563t.137-1.062l6.9-8.425q.3-.375.775-.375t.775.375l6.9 8.425q.4.5.138 1.063T18.9 13H16v4q0 .425-.288.713T15 18H9q-.425 0-.713-.288T8 17Zm2-1h4v-5h2.775L12 5.15L7.225 11H10v5Zm2-5.425ZM5 22q-.425 0-.713-.288T4 21q0-.425.288-.713T5 20h14q.425 0 .713.288T20 21q0 .425-.288.713T19 22H5Z"/>`),
		g.Group(children),
	)
}