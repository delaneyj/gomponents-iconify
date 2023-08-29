package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeavePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><rect width="10" height="15" x="6" y="4" fill="currentColor" opacity=".8" rx="1"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.667 7.5L17.75 10l-2.083-2.5Zm0 5L17.75 10l-2.083 2.5Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" d="M17 10H9.5M4 3h9M4 17h9m0-14v4m0 6v4M4 3v14"/></g>`),
		g.Group(children),
	)
}