package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserTalkRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 2v12c0 1.1.9 2 2 2h14l4 4V2c0-1.1-.9-2-2-2H2C.9 0 0 .9 0 2zm7.5 3.5C7.5 6.3 6.8 7 6 7s-1.5-.7-1.5-1.5S5.2 4 6 4s1.5.7 1.5 1.5zm8 0c0 .8-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5S13.2 4 14 4s1.5.7 1.5 1.5zM4.4 9h11.3c-.8 2.3-3 3-5.6 3s-4.9-.7-5.7-3z"/>`),
		g.Group(children),
	)
}