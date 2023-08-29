package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FurAffinityNoc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="23.915" cy="32.457" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="11.464" ry="9.129"/><ellipse cx="23.915" cy="13.861" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.449" ry="7.447"/><ellipse cx="37.039" cy="17.585" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.066" ry="5.449" transform="rotate(-85.419 37.039 17.585)"/><ellipse cx="10.961" cy="17.585" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.449" ry="7.066" transform="rotate(-4.581 10.961 17.584)"/>`),
		g.Group(children),
	)
}