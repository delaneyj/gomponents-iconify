package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentArchiveBoxContentBankerArchiveFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 5.54h11v7a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-7h0Z"/><rect width="4" height="13" x="5" y="-2.96" rx="1" transform="rotate(90 7 3.54)"/><path d="M5.5 8.54h3"/></g>`),
		g.Group(children),
	)
}