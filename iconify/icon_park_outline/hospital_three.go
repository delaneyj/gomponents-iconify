package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M11 26v18h26V26L24 13L11 26Z"/><path fill="currentColor" d="M3.586 21.586a2 2 0 1 0 2.828 2.828l-2.828-2.828ZM24 4l1.414-1.414a2 2 0 0 0-2.828 0L24 4Zm17.586 20.414a2 2 0 1 0 2.828-2.828l-2.828 2.828Zm-35.172 0l19-19l-2.828-2.828l-19 19l2.828 2.828Zm16.172-19l19 19l2.828-2.828l-19-19l-2.828 2.828Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 31h10m-5-5v10"/></g>`),
		g.Group(children),
	)
}