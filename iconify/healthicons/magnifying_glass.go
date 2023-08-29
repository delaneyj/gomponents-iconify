package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M18.748 12.816c-1.74.067-3.313.688-4.154 1.53a1 1 0 1 1-1.414-1.415c1.297-1.297 3.409-2.033 5.49-2.114c2.095-.081 4.382.492 5.984 2.094a1 1 0 0 1-1.415 1.414c-1.09-1.091-2.764-1.577-4.491-1.51Z"/><path fill-rule="evenodd" d="M28.936 27.384A12.948 12.948 0 0 0 32 19c0-7.18-5.82-13-13-13S6 11.82 6 19s5.82 13 13 13c3.195 0 6.12-1.152 8.384-3.064L27 31.144l7.514 7.514l4.144-4.144L31.144 27l-2.208.384ZM19 29c5.523 0 10-4.477 10-10S24.523 9 19 9S9 13.477 9 19s4.477 10 10 10Z" clip-rule="evenodd"/><path d="m35.928 40.072l4.144-4.144l1.356 1.356c.763.763.763 2 0 2.762l-1.382 1.382c-.763.763-2 .763-2.762 0l-1.356-1.356Z"/></g>`),
		g.Group(children),
	)
}