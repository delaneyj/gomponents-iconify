package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAwardTrophyRewardRatingTrophySocialAwardMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 8.5v5m-2.5 0h5m-6-8a3 3 0 0 1-3-3v-1H4v4Zm7 0a3 3 0 0 0 3-3v-1H10v4Z"/><path d="M10 5.5a3 3 0 0 1-6 0v-5h6Z"/></g>`),
		g.Group(children),
	)
}