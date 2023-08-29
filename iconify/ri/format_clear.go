package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.651 14.065L11.605 20H9.574l1.35-7.661l-7.41-7.41L4.93 3.515L20.485 19.07l-1.414 1.414l-6.42-6.42Zm-.878-6.535l.27-1.53h-1.8l-2-2H20v2h-5.927L13.5 9.256L11.773 7.53Z"/>`),
		g.Group(children),
	)
}