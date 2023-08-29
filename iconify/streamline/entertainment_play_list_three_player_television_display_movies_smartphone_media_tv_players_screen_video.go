package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentPlayListThreePlayerTelevisionDisplayMoviesSmartphoneMediaTvPlayersScreenVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8.5" height="13" x="2.75" y=".5" rx="1"/><path d="M5.27 7.14V3.82a.31.31 0 0 1 .47-.28l2.85 1.67a.31.31 0 0 1 0 .54L5.74 7.41a.31.31 0 0 1-.47-.27ZM2.75 10.5h8.5"/></g>`),
		g.Group(children),
	)
}