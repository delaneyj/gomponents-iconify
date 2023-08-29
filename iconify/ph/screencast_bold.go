package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 60v136a20 20 0 0 1-20 20h-40a12 12 0 0 1 0-24h36V64H48v4a12 12 0 0 1-24 0v-8a20 20 0 0 1 20-20h168a20 20 0 0 1 20 20ZM36 176a12 12 0 0 0 0 24a4 4 0 0 1 4 4a12 12 0 0 0 24 0a28 28 0 0 0-28-28Zm0-40a12 12 0 0 0 0 24a44.05 44.05 0 0 1 44 44a12 12 0 0 0 24 0a68.07 68.07 0 0 0-68-68Zm0-40a12 12 0 0 0 0 24a84.09 84.09 0 0 1 84 84a12 12 0 0 0 24 0A108.12 108.12 0 0 0 36 96Z"/>`),
		g.Group(children),
	)
}