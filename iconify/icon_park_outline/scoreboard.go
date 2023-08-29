package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scoreboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="12" rx="3"/><path stroke-linecap="round" d="M14 6v6m20-6v6M10.227 24L15 19.017V33m9-21v28"/><ellipse cx="34" cy="26" rx="3" ry="7"/><path stroke-linecap="round" d="M21 12h6m-6 28h6"/></g>`),
		g.Group(children),
	)
}