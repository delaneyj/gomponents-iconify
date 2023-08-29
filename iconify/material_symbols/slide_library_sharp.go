package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideLibrarySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V7h2v11h14v2H2Zm4-4V4h9.4l-2 2H8v8h12V8.6l2-2V16H6Zm7.5-4q-.575 0-1.113-.263T11.5 11q.425 0 .713-.288T12.5 10q0-.625.438-1.063T14 8.5q.625 0 1.063.438T15.5 10q0 .825-.588 1.413T13.5 12Zm3.275-3L15 7.225L19.45 2.8l1.75 1.75L16.775 9Z"/>`),
		g.Group(children),
	)
}