package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RankingList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRankingList0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 8H6a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2Z"/><path stroke="#000" stroke-linecap="round" d="M24 17v14m8-7v7m-16-9v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRankingList0)"/>`),
		g.Group(children),
	)
}