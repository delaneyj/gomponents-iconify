package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendTimeExtensionOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22v-4l4-1l-4-1v-4l10 5Zm-8-1q-.825 0-1.413-.587Q3 19.825 3 19v-3.8q1.2 0 2.1-.762q.9-.763.9-1.938q0-1.175-.9-1.938Q4.2 9.8 3 9.8V6q0-.825.587-1.412Q4.175 4 5 4h4q0-1.05.725-1.775Q10.45 1.5 11.5 1.5q1.05 0 1.775.725Q14 2.95 14 4h4q.825 0 1.413.588Q20 5.175 20 6v7.25l-2-1V6H5v2.2q1.35.5 2.175 1.675Q8 11.05 8 12.5q0 1.425-.825 2.6T5 16.8V19h2.125q.425-1.125 1.45-1.962Q9.6 16.2 11 16.05v2q-1 .2-1.6.938q-.6.737-.6 2.012Zm6.5-9.75Z"/>`),
		g.Group(children),
	)
}