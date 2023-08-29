package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 128v80a12 12 0 0 1-12 12H48a12 12 0 0 1-12-12V56a12 12 0 0 1 12-12h80a4 4 0 0 1 0 8H48a4 4 0 0 0-4 4v152a4 4 0 0 0 4 4h152a4 4 0 0 0 4-4v-80a4 4 0 0 1 8 0Zm16-68a32 32 0 1 1-32-32a32 32 0 0 1 32 32Zm-8 0a24 24 0 1 0-24 24a24 24 0 0 0 24-24Z"/>`),
		g.Group(children),
	)
}