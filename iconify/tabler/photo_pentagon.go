package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoPentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.163 2.168l8.021 5.828c.694.504.984 1.397.719 2.212l-3.064 9.43a1.978 1.978 0 0 1-1.881 1.367H7.042a1.978 1.978 0 0 1-1.881-1.367l-3.064-9.43a1.978 1.978 0 0 1 .719-2.212l8.021-5.828a1.978 1.978 0 0 1 2.326 0zM15 8h.01"/><path d="m4 15l4-4c.928-.893 2.072-.893 3 0l5 5"/><path d="m14 14l1-1c.928-.893 2.072-.893 3 0l2 2"/></g>`),
		g.Group(children),
	)
}