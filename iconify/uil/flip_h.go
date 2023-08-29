package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 11H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2ZM10.93 9h1.5a1 1 0 0 0 0-2h-1.5a1 1 0 0 0 0 2Zm4.5-1a1 1 0 0 0 1 1H17a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09l-.66-.65a1 1 0 0 0-1.41 0a1 1 0 0 0-.19 1.15a1.49 1.49 0 0 0-.02.21Zm-3.78-3.23l.35-.36l.81.81a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41l-1.06-1.06a.91.91 0 0 0-.26-.19a1 1 0 0 0-1.61-.27l-1.06 1.06a1 1 0 0 0 1.42 1.42ZM17 15H7a1 1 0 0 0-.92.62a1 1 0 0 0 .21 1.09l5 5a1 1 0 0 0 1.42 0l5-5a1 1 0 0 0 .21-1.09A1 1 0 0 0 17 15Zm-5 4.59L9.41 17h5.18ZM7.05 9a1 1 0 0 0 .71-.29L8.82 7.6A1 1 0 0 0 7.4 6.18L6.34 7.24a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.34Z"/>`),
		g.Group(children),
	)
}