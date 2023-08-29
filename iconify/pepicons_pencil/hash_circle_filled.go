package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilHashCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 10.75a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5Zm0 5.5a.5.5 0 0 1 .5-.5h13a.5.5 0 1 1 0 1h-13a.5.5 0 0 1-.5-.5Z"/><path d="M10.791 6.502a.5.5 0 0 1 .457.54l-1 12a.5.5 0 0 1-.996-.084l1-12a.5.5 0 0 1 .54-.456Zm5.5 0a.5.5 0 0 1 .457.54l-1 12a.5.5 0 0 1-.996-.084l1-12a.5.5 0 0 1 .54-.456Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilHashCircleFilled0)"/></g>`),
		g.Group(children),
	)
}