package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BillBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.25a.75.75 0 0 0 0 1.5h2v9.527c0 1.34 0 2.01.268 2.601c.268.591.772 1.032 1.781 1.915l2 1.75c1.883 1.647 2.824 2.47 3.951 2.47c1.127 0 2.069-.823 3.951-2.47l2-1.75c1.008-.883 1.513-1.324 1.78-1.915c.269-.59.269-1.26.269-2.6V2.75h2a.75.75 0 0 0 0-1.5H2Zm6.5 11a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5h-7ZM7.75 8a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5h-7A.75.75 0 0 1 7.75 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}