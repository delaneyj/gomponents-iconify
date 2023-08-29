package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFourBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m230.86 109.25l-61.68-22.43l-22.43-61.68a19.95 19.95 0 0 0-37.5 0L86.82 86.82l-61.68 22.43a19.95 19.95 0 0 0 0 37.5l61.68 22.43l22.43 61.68a19.95 19.95 0 0 0 37.5 0l22.43-61.68l61.68-22.43a19.95 19.95 0 0 0 0-37.5Zm-71.65 38a19.92 19.92 0 0 0-11.94 11.94l-19.27 53l-19.27-53a19.92 19.92 0 0 0-11.94-11.94l-53-19.25l53-19.27a19.92 19.92 0 0 0 11.94-11.94l19.27-53l19.27 53a19.92 19.92 0 0 0 11.94 11.94l53 19.27Z"/>`),
		g.Group(children),
	)
}