package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapIdentify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 8h2v6h-2zm0 10h2v6h-2zm3-3h6v2h-6zM8 15h6v2H8zm-4-5H2V4a2.002 2.002 0 0 1 2-2h6v2H4zm6 20H4a2.002 2.002 0 0 1-2-2v-6h2v6h6zm18 0h-6v-2h6v-6h2v6a2.002 2.002 0 0 1-2 2zm2-20h-2V4h-6V2h6a2.002 2.002 0 0 1 2 2z"/>`),
		g.Group(children),
	)
}