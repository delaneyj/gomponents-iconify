package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmpleDress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m288 16l32 16s-25.2 44.02-16 64c5 10.8 32 16 32 16c-16 32-32 80-32 96c80 48 80 144 160 176c0 64-80 112-208 112S48 448 48 384c80-32 80-128 160-176c0-16-16-64-32-96c0 0 27-5.2 32-16c9.2-19.98-16-64-16-64l32-16c0 32 16 48 32 48s32-16 32-48z"/>`),
		g.Group(children),
	)
}