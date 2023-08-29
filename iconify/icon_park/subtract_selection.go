package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="27" height="27" x="16" y="16" rx="2"/><path d="M16 32H7C5.89543 32 5 31.1046 5 30V7C5 5.89543 5.89543 5 7 5H30C31.1046 5 32 5.89543 32 7V16"/><path d="M29 16L16 30"/><path d="M38 16L16 40"/><path d="M43 21L23 43"/><path d="M43 32L33 43"/></g>`),
		g.Group(children),
	)
}