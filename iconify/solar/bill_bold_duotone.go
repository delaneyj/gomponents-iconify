package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BillBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.25 3A.75.75 0 0 1 2 2.25h20a.75.75 0 0 1 0 1.5H2A.75.75 0 0 1 1.25 3Z" clip-rule="evenodd"/><path d="M4 3.75v9.527c0 1.34 0 2.01.268 2.601c.268.591.772 1.032 1.781 1.915l2 1.75c1.883 1.647 2.824 2.47 3.951 2.47c1.127 0 2.069-.823 3.951-2.47l2-1.75c1.008-.883 1.513-1.324 1.78-1.915c.269-.59.269-1.26.269-2.6V3.75H4Z" opacity=".5"/><path fill-rule="evenodd" d="M7.75 14a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5h-7a.75.75 0 0 1-.75-.75Zm0-5a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5h-7A.75.75 0 0 1 7.75 9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}