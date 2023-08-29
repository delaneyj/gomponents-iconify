package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactPageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14Zm-4 4h8v-.575q0-.6-.325-1.1t-.9-.75q-.65-.275-1.35-.425T12 15q-.725 0-1.425.15t-1.35.425q-.575.25-.9.75T8 17.425V18Zm12 4H4V2h10l6 6v14Z"/>`),
		g.Group(children),
	)
}