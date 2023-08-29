package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockersBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 80a12 12 0 0 1-12 12h-16a12 12 0 0 1 0-24h16a12 12 0 0 1 12 12Zm-12 28h-16a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24ZM76 92h16a12 12 0 0 0 0-24H76a12 12 0 0 0 0 24Zm16 16H76a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24Zm136-60v176a12 12 0 0 1-24 0v-12h-64v12a12 12 0 0 1-24 0v-12H52v12a12 12 0 0 1-24 0V48a20 20 0 0 1 20-20h160a20 20 0 0 1 20 20ZM116 188V52H52v136Zm24 0h64V52h-64Z"/>`),
		g.Group(children),
	)
}