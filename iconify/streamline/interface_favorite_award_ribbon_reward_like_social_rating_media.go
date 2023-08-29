package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFavoriteAwardRibbonRewardLikeSocialRatingMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.87" cy="5" r="4.5"/><circle cx="6.87" cy="5" r="2"/><path d="m6 9.42l-.88 3.7a.51.51 0 0 1-.26.33a.54.54 0 0 1-.43 0L1.11 12a.51.51 0 0 1-.18-.78L3.5 8M8 9.37l.9 3.75a.5.5 0 0 0 .27.33a.51.51 0 0 0 .42 0l3.3-1.45a.5.5 0 0 0 .28-.35a.48.48 0 0 0-.1-.43l-2.68-3.41"/></g>`),
		g.Group(children),
	)
}