package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11.075 37.08c-3.386 1.46-5.039 3.444-5.71 4.83c-.266.548.181 1.09.79 1.09h21.232c.742 0 1.212-.776.779-1.38c-1.75-2.438-4.384-3.977-5.777-4.552a.933.933 0 0 0-.357-.068h-10.57a.98.98 0 0 0-.387.08Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m14.732 18.856l5-8.66l8.66 5l-5 8.66z"/><path fill="currentColor" d="m27.392 16.928l1-1.732a2 2 0 0 0-2.732.732l1.732 1Zm-3 5.196l-1.732-1a2 2 0 0 0 .732 2.732l1-1.732Zm2-3.464l12.99 7.5l2-3.464l-12.99-7.5l-2 3.464Zm11.99 9.232l-12.99-7.5l-2 3.464l12.99 7.5l2-3.464Zm-12.258-4.768l3-5.196l-3.464-2l-3 5.196l3.464 2Zm13.625 4.402a1 1 0 0 1-1.367.366l-2 3.464a5 5 0 0 0 6.83-1.83l-3.464-2Zm-.367-1.366a1 1 0 0 1 .367 1.366l3.464 2a5 5 0 0 0-1.83-6.83l-2 3.464Z"/><rect width="14" height="6" x="21" y="4" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="3" transform="rotate(30 21 4)"/><rect width="14" height="6" x="13" y="17.856" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="3" transform="rotate(30 13 17.856)"/></g>`),
		g.Group(children),
	)
}