package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxMultipleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 7V3a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-4v3.992C17 21.55 16.551 22 15.992 22H3.008A1.006 1.006 0 0 1 2 20.992l.003-12.985C2.003 7.451 2.452 7 3.01 7H7Zm2 0h6.993C16.549 7 17 7.449 17 8.007V15h3V4H9v3Zm6 2H4.003L4 20h11V9Zm-6.498 9l-3.535-3.536L6.38 13.05l2.121 2.122l4.243-4.243l1.414 1.414L8.502 18Z"/>`),
		g.Group(children),
	)
}