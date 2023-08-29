package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackdiamondsuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FF473E" d="M248.07 502.699L86.126 261.32a9.549 9.549 0 0 1 0-10.64L248.07 9.301c3.783-5.639 12.077-5.639 15.859 0L425.874 250.68a9.549 9.549 0 0 1 0 10.64L263.93 502.699c-3.783 5.639-12.077 5.639-15.86 0z"/>`),
		g.Group(children),
	)
}