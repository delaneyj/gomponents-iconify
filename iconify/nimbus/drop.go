package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 12.09a1.74 1.74 0 0 1-1.25-.49a1.1 1.1 0 0 1-.36-.8a.63.63 0 0 0-.63-.63a.62.62 0 0 0-.62.63a2.33 2.33 0 0 0 .73 1.69a3 3 0 0 0 2.13.85a.63.63 0 1 0 0-1.25z"/><path fill="currentColor" d="M9.25 1.25C9.06.23 8.43 0 8 0S6.94.23 6.75 1.25C6 5.2 2.5 7 2.5 10.82A5.35 5.35 0 0 0 8 16a5.35 5.35 0 0 0 5.5-5.18C13.5 7 10 5.2 9.25 1.25zM8 14.75a4.1 4.1 0 0 1-4.25-3.93c0-1.66.85-2.9 1.84-4.33A12.85 12.85 0 0 0 8 1.48v-.1v.1a12.85 12.85 0 0 0 2.39 5c1 1.43 1.84 2.67 1.84 4.33A4.1 4.1 0 0 1 8 14.75z"/>`),
		g.Group(children),
	)
}