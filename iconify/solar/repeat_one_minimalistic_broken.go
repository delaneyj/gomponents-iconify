package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOneMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M16 12H8"/><path fill="currentColor" d="M9.5 19.75a.75.75 0 0 0 0-1.5v1.5ZM11 5v.75a.75.75 0 0 0 .53-1.28L11 5ZM9.53 2.47a.75.75 0 0 0-1.06 1.06l1.06-1.06ZM1.25 12a.75.75 0 0 0 1.5 0h-1.5Zm2.62 3.57a.75.75 0 0 0-1.23.86l1.23-.86Zm5.63 2.68H9v1.5h.5v-1.5ZM9 5.75h2v-1.5H9v1.5Zm2.53-1.28l-2-2l-1.06 1.06l2 2l1.06-1.06ZM2.75 12A6.25 6.25 0 0 1 9 5.75v-1.5A7.75 7.75 0 0 0 1.25 12h1.5Zm-.11 4.43A7.743 7.743 0 0 0 9 19.75v-1.5a6.24 6.24 0 0 1-5.13-2.68l-1.23.86ZM13 19v-.75a.75.75 0 0 0-.53 1.28L13 19Zm1.47 2.53a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm.03-17.28a.75.75 0 0 0 0 1.5v-1.5ZM22.75 12a.75.75 0 0 0-1.5 0h1.5Zm-2.62-3.57a.75.75 0 1 0 1.23-.86l-1.23.86ZM15 18.25h-2v1.5h2v-1.5Zm-2.53 1.28l2 2l1.06-1.06l-2-2l-1.06 1.06ZM14.5 5.75h.5v-1.5h-.5v1.5ZM21.25 12A6.25 6.25 0 0 1 15 18.25v1.5A7.75 7.75 0 0 0 22.75 12h-1.5Zm.11-4.43A7.742 7.742 0 0 0 15 4.25v1.5c2.123 0 4 1.058 5.13 2.68l1.23-.86Z"/></g>`),
		g.Group(children),
	)
}