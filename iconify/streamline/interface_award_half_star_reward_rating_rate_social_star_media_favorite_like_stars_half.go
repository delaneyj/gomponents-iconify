package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAwardHalfStarRewardRatingRateSocialStarMediaFavoriteLikeStarsHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25.5a.54.54 0 0 0-.49.32L8.17 4.18a.52.52 0 0 1-.41.31L4.22 5a.58.58 0 0 0-.3 1l2.56 2.63a.58.58 0 0 1 .16.5L6 12.83a.56.56 0 0 0 .8.6l3.2-1.74a.59.59 0 0 1 .26-.07Z"/>`),
		g.Group(children),
	)
}