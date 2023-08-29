package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSpecialRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm10.9-5.05l1.7 1.275q.15.125.288.025t.087-.275l-.625-2.125l1.75-1.4q.125-.125.075-.288T17.95 12H15.8l-.65-2.05q-.05-.175-.25-.175t-.25.175L14 12h-2.15q-.175 0-.225.162t.075.288l1.75 1.4l-.625 2.125q-.05.175.087.275t.288-.025l1.7-1.275Z"/>`),
		g.Group(children),
	)
}