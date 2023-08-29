package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChannelTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.168a1.75 1.75 0 1 0 1.5 3.163A1.75 1.75 0 0 0 3 7.168Zm.168-1.106A2.75 2.75 0 1 1 3 11.397v9.353A4.25 4.25 0 0 0 7.25 25h13.5A4.25 4.25 0 0 0 25 20.75V7.25A4.25 4.25 0 0 0 20.75 3H7.25a4.252 4.252 0 0 0-4.082 3.062ZM9.75 11h8.5a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1 0-1.5ZM9 15.75a.75.75 0 0 1 .75-.75h5.75a.75.75 0 0 1 0 1.5H9.75a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}