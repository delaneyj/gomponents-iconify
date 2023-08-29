package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M29.799 27.43a1.25 1.25 0 0 1-1.049.57h-10.5c-.69 0-1.25-.56-1.25-1.25V3.25a1.25 1.25 0 0 1 2.391-.51l10.5 23.5a1.25 1.25 0 0 1-.092 1.19ZM19.5 9.112V25.5h7.322L19.5 9.112ZM3 28a1 1 0 0 1-.91-1.416l11-24A1 1 0 0 1 15 3v24a1 1 0 0 1-1 1H3Z"/>`),
		g.Group(children),
	)
}