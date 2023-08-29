package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KipSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M340.8 88.3c13.4-11.5 15-31.7 3.5-45.1s-31.7-15-45.1-3.5L128 186.4V64c0-17.7-14.3-32-32-32S64 46.3 64 64v160H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h32v160c0 17.7 14.3 32 32 32s32-14.3 32-32V325.6l171.2 146.7c13.4 11.5 33.6 9.9 45.1-3.5s9.9-33.6-3.5-45.1L182.5 288H352c17.7 0 32-14.3 32-32s-14.3-32-32-32H182.5L340.8 88.3z"/>`),
		g.Group(children),
	)
}