package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FerryJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9.5 12.5h-4V8c0-1.088.58-2.552 1.354-3.862c.22-.371.44-.708.646-.993c.205.285.427.622.646.993C8.92 5.448 9.5 6.912 9.5 8v4.5ZM7.5 1C7 1 4 5 4 8v4.5A1.5 1.5 0 0 0 5.5 14h4a1.5 1.5 0 0 0 1.5-1.5V8c0-3-3-7-3.5-7Z"/>`),
		g.Group(children),
	)
}