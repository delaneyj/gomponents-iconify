package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InnerShadowLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M4.929 4.929c3.905-3.905 10.237-3.905 14.142 0c3.905 3.905 3.905 10.237 0 14.142c-3.905 3.905-10.237 3.905-14.142 0c-3.905-3.905-3.905-10.237 0-14.142zM8.464 7.05a1 1 0 0 0-1.414 0a7 7 0 0 0 0 9.9a1 1 0 1 0 1.414-1.414a5 5 0 0 1 0-7.072a1 1 0 0 0 0-1.414z"/></g>`),
		g.Group(children),
	)
}