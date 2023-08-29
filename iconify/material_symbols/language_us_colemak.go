package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageUsColemak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15H18V9h-3.5v6ZM6 17q-.825 0-1.413-.588T4 15V9q0-.825.588-1.413T6 7h4.5v2H6v6h4.5v2H6Zm8.5 0q-.825 0-1.413-.588T12.5 15V9q0-.825.588-1.413T14.5 7H18q.825 0 1.413.588T20 9v6q0 .825-.588 1.413T18 17h-3.5Z"/>`),
		g.Group(children),
	)
}