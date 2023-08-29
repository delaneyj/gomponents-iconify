package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upgrade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 29.816l-231 154v106.368l231-154l231 154V183.816zm0 128.043L105 259.783v90.283l151-101.925l151 101.925v-90.283zm0 112l-87 58.725v67.6l87-58l87 58v-67.6zm0 89.957l-87 58v64.368l87-58l87 58v-64.368z"/>`),
		g.Group(children),
	)
}