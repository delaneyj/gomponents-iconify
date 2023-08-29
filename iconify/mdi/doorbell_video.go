package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorbellVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 15c0 1.11-.89 2-2 2s-2-.89-2-2s.9-2 2-2s2 .9 2 2m4-11v16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4c0-1.1.9-2 2-2h8a2 2 0 0 1 2 2m-7.5 3c0 .83.67 1.5 1.5 1.5s1.5-.67 1.5-1.5s-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5m5.5 3H8v10h8V10Z"/>`),
		g.Group(children),
	)
}