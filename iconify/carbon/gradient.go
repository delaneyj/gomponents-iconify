package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gradient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 4H6a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2Zm-4 22v-4h-4v4h-4v-4h-4v-4h4v-4h-4v-4h4V6h4v4h4V6h4v20Z"/><path fill="currentColor" d="M14 10h4v4h-4zm0 8h4v4h-4zm4-4h4v4h-4z"/>`),
		g.Group(children),
	)
}