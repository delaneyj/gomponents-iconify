package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subtract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSubtract0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSubtract1" fill="currentColor"><path id="feSubtract2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-4-2h8V5h-8v8Z"/></g></g>`),
		g.Group(children),
	)
}