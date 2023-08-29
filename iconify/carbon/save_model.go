package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveModel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26 20l1.427 1.903L23 26.963l-4.427-5.06L20 20h6m1-2h-8l-3 4l7 8l7-8l-3-4Z"/><path fill="currentColor" d="M16 26h-4v-8h2v-2h-2a2 2 0 0 0-2 2v8H6V6h4v4a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V6.41l4 4V16h2v-6a1 1 0 0 0-.29-.71l-5-5A.999.999 0 0 0 22 4H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h10ZM12 6h8v4h-8Z"/>`),
		g.Group(children),
	)
}