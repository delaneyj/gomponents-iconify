package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMusicList0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 19h16m-16-9h16M8 38h32M8 28h32"/><path fill="#fff" d="m8 10l8 5l-8 5V10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMusicList0)"/>`),
		g.Group(children),
	)
}