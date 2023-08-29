package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CannedFruit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M15 24s-.5 3 1 9"/><path stroke-linecap="round" stroke-linejoin="round" d="M37.567 12C38.727 14.81 40 18.642 40 21.938c0 3.877-1.345 11.412-2.315 16.339c-.544 2.765-2.982 4.723-5.8 4.723H16.042c-2.785 0-5.203-1.914-5.762-4.642C9.314 33.638 8 26.402 8 21.938c0-3.679 1.444-7.306 2.827-9.938"/><path d="M11.587 6.457a1 1 0 0 1 .84-.457h23.147a1 1 0 0 1 .84.457l2.587 4a1 1 0 0 1-.84 1.543H9.839a1 1 0 0 1-.84-1.543l2.589-4Z"/></g>`),
		g.Group(children),
	)
}