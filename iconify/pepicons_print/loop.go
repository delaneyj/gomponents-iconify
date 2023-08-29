package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.5 9.5a5.5 5.5 0 0 0 8.22 4.782c.06.099.133.193.22.279l3 3a1.5 1.5 0 0 0 2.12-2.122l-3-3a1.505 1.505 0 0 0-.278-.22A5.5 5.5 0 1 0 5.5 9.5Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M4.828 4.828A5 5 0 1 0 11.9 11.9a5 5 0 0 0-7.07-7.07Zm6.364 6.364a4 4 0 1 1-5.656-5.657a4 4 0 0 1 5.656 5.657Z" clip-rule="evenodd"/><path d="M11.192 12.627a1 1 0 0 1 1.415-1.414l2.828 2.829a1 1 0 1 1-1.414 1.414l-2.829-2.829Z"/></g>`),
		g.Group(children),
	)
}