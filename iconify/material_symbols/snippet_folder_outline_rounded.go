package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnippetFolderOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm0-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Zm10.5-2.5v-5h1.375l1.625 1.625V15.5h-3ZM14 17h4q.425 0 .713-.288T19 16v-4.075q0-.2-.075-.388T18.7 11.2l-1.9-1.9q-.15-.15-.338-.225T16.076 9H14q-.425 0-.712.288T13 10v6q0 .425.288.713T14 17Z"/>`),
		g.Group(children),
	)
}