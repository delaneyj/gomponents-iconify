package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentPlayListOneScreenTelevisionDisplayPlayerMoviesMovieTvMediaPlayersVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M5.49 10.56V6.73A.36.36 0 0 1 6 6.42l3.32 1.91a.37.37 0 0 1 0 .63L6 10.88a.37.37 0 0 1-.51-.32ZM.5 4h13M4 4L5.5.5m3 3.5L10 .5"/></g>`),
		g.Group(children),
	)
}