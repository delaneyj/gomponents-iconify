package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 5q18 0 30.5 12.5T384 48v276q0 23-19 35L192 475L19 359Q0 347 0 324V48q0-18 12.5-30.5T43 5h298zM149 325l192-192l-30-30l-162 162l-76-76l-30 30z"/>`),
		g.Group(children),
	)
}