package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 12H4V4h8v2H6v6zm22 0h-2V6h-6V4h8v8zM12 28H4v-8h2v6h6v2zm16 0h-8v-2h6v-6h2v8zM15 10h2v4h-2zm-5 5h4v2h-4zm8 0h4v2h-4zm-3 3h2v4h-2z"/>`),
		g.Group(children),
	)
}