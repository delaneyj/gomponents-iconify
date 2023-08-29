package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InnerShadowBottomFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M5.144 4.72c3.92-3.695 10.093-3.625 13.927.209c3.905 3.905 3.905 10.237 0 14.142c-3.905 3.905-10.237 3.905-14.142 0c-3.905-3.905-3.905-10.237 0-14.142zm3.32 10.816A1 1 0 1 0 7.05 16.95a7 7 0 0 0 9.9 0a1 1 0 0 0-1.414-1.414a5 5 0 0 1-7.072 0z"/></g>`),
		g.Group(children),
	)
}