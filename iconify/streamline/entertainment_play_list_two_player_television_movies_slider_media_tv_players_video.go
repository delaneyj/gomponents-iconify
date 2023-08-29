package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentPlayListTwoPlayerTelevisionMoviesSliderMediaTvPlayersVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 8.68V5.32c0-.25.23-.4.41-.28l2.45 1.68a.35.35 0 0 1 0 .56L5.91 9c-.18.08-.41-.07-.41-.32Z"/><rect width="8" height="11" x="3" y="1.5" rx="1"/><path d="M.5 3v8m13-8v8"/></g>`),
		g.Group(children),
	)
}