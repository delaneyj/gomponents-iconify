package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioVideoReceiverSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V5h20v12h-2v2h-2v-2H6v2H4v-2H2Zm14.5-4q.825 0 1.413-.588T18.5 11q0-.825-.588-1.413T16.5 9q-.825 0-1.413.588T14.5 11q0 .825.588 1.413T16.5 13ZM6 13h7V9H6v4Z"/>`),
		g.Group(children),
	)
}