package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.719 5.281L5.28 6.72l20 20l1.438-1.438L21.437 20H27v-2h-7.563l-4-4H27v-2H13.437zM5 12v2h4.906l-2-2zm0 6v2h10.906l-2-2z"/>`),
		g.Group(children),
	)
}