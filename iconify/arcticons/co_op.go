package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoOp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="7.5" height="9.938" x="26.146" y="12" rx="3.75" ry="3.75"/><rect width="7.5" height="9.938" x="14.354" y="25" rx="3.75" ry="3.75"/><path d="M21.36 20.049a3.749 3.749 0 0 1-3.256 1.889h0a3.75 3.75 0 0 1-3.75-3.75V15.75a3.75 3.75 0 0 1 3.75-3.75h0c1.39 0 2.605.757 3.252 1.882M26.146 40V28.75a3.75 3.75 0 0 1 3.75-3.75h0a3.75 3.75 0 0 1 3.75 3.75v2.438a3.75 3.75 0 0 1-3.75 3.75H28.33"/></g>`),
		g.Group(children),
	)
}