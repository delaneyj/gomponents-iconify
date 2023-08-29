package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatsCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236.06 187.69A84 84 0 0 0 172.29 68.9a84 84 0 1 0-152.35 70.79l-7.24 25.36A18 18 0 0 0 35 187.3l25.36-7.24a84.27 84.27 0 0 0 23.36 7a84.05 84.05 0 0 0 112 41l25.36 7.24a18 18 0 0 0 22.25-22.25Zm-207.7-23.47Zm33.17-9a12 12 0 0 0-3.3.46l-19.49 5.57l5.57-19.49a12 12 0 0 0-1-9.05a60 60 0 1 1 24 24a11.91 11.91 0 0 0-5.78-1.48Zm150.16 34.54l5.57 19.49l-19.49-5.57a12 12 0 0 0-9.05 1A60.06 60.06 0 0 1 111 186.63a83.93 83.93 0 0 0 68.55-91.37a60 60 0 0 1 33.16 85.46a12 12 0 0 0-1.02 9.05Z"/>`),
		g.Group(children),
	)
}