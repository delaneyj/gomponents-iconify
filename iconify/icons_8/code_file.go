package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8V5zm12 1.438L22.563 9H20V6.437zM16 13l-2 12h2l2-12h-2zm-3.78 2.375l-2.5 3l-.533.625l.532.625l2.5 3l1.56-1.25L11.813 19l1.968-2.375l-1.56-1.25zm7.56 0l-1.56 1.25L20.187 19l-1.97 2.375l1.563 1.25l2.5-3l.532-.625l-.53-.625l-2.5-3z"/>`),
		g.Group(children),
	)
}