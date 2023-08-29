package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wheat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 3v2.586l-4 4V1h-2v8.586l-4-4V3H9v16a7.004 7.004 0 0 0 6 6.92V30h2v-4.08A7.004 7.004 0 0 0 23 19V3Zm-6 20.899A5.008 5.008 0 0 1 11 19v-2.586l4 4Zm0-6.313l-4-4V8.414l4 4Zm2-5.172l4-4v5.172l-4 4Zm0 11.485v-3.485l4-4V19a5.008 5.008 0 0 1-4 4.899Z"/>`),
		g.Group(children),
	)
}