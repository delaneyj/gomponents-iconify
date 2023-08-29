package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccumulationIce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 8v8l1 2l1-2V8h-2zm-4 0v10l1 2l1-2V8h-2z"/><path fill="currentColor" d="M28 4a2.002 2.002 0 0 0-2 2v20H6v-4h4v-2H6v-4h4v-2H6v-4h8v4l1 2l1-2V8H6V6a2.002 2.002 0 0 0-2-2H2v2h2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V6h2V4Z"/>`),
		g.Group(children),
	)
}