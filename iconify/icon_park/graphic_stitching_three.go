package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicStitchingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M24 14C26.7614 14 29 11.7614 29 9C29 6.23858 26.7614 4 24 4C21.2386 4 19 6.23858 19 9C19 11.7614 21.2386 14 24 14Z"/><path fill="#2F88FF" d="M24 44C26.7614 44 29 41.7614 29 39C29 36.2386 26.7614 34 24 34C21.2386 34 19 36.2386 19 39C19 41.7614 21.2386 44 24 44Z"/><path fill="#2F88FF" d="M14 19H4V29H14V19Z"/><path fill="#2F88FF" d="M44 19H34V29H44V19Z"/><path d="M19 9H9V19"/><path d="M19 39H9V29"/><path d="M29 9H40V19"/><path d="M29 39H39V29"/></g>`),
		g.Group(children),
	)
}