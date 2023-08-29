package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiooff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M4.5 29.479h8l9 11.015c1 1.44 3.34 1.331 4-.302V2.83c-.811-1.632-2.939-1.763-4-.382l-9 11.053h-8c-2.561 0-3 .461-3 2.964v10.012c0 2.442.5 3.002 3 3.002z"/>`),
		g.Group(children),
	)
}