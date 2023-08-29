package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 8a1.5 1.5 0 0 1 1.5-1.5h10a1.5 1.5 0 0 1 0 3h-10A1.5 1.5 0 0 1 5 8Zm0 6a1.5 1.5 0 0 1 1.5-1.5h10a1.5 1.5 0 0 1 0 3h-10A1.5 1.5 0 0 1 5 14Z" opacity=".2"/><path d="M4.5 7a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Z"/></g>`),
		g.Group(children),
	)
}