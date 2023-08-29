package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 15a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z" opacity=".16"/><path d="M13.847 4.532a1 1 0 1 0-1.694-1.064l1.694 1.064Zm-6.92 7.264A1 1 0 0 0 8.62 12.86l-1.694-1.064ZM16 15a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6h-2Zm-4 4a4 4 0 0 1-4-4H6a6 6 0 0 0 6 6v-2Zm-4-4a4 4 0 0 1 4-4V9a6 6 0 0 0-6 6h2Zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6v2Zm.153-7.532l-5.227 8.328L8.62 12.86l5.227-8.328l-1.694-1.064Z"/></g>`),
		g.Group(children),
	)
}