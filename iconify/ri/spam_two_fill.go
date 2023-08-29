package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpamTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.935 2.501L21.5 8.066v7.87l-5.565 5.565h-7.87L2.5 15.936v-7.87l5.565-5.565h7.87Zm-4.936 12.5v2h2v-2h-2Zm0-8v6h2V7h-2Z"/>`),
		g.Group(children),
	)
}