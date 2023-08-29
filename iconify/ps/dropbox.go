package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M375 161L238 77l87-72l137 85zm-51 159l132-78l-80-67l-129 78zm-107-67L88 175L8 242l132 78zM94 309v25l132 81V261l-87 72zm276 0l-45 24l-87-72v154l132-81v-25zM226 77L140 5L2 90l87 71z"/>`),
		g.Group(children),
	)
}