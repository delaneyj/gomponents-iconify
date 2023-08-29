package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSymlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 13.25l3.5-3.5m0 3v-3h-3"/><path d="M8.25 13.25h6v-8.5h-6l-1.5-2h-5v4"/></g>`),
		g.Group(children),
	)
}