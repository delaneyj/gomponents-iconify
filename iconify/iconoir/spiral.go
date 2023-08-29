package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spiral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.5 6.004C3.5 7.808 6.357 9 11.5 9c7 0 8-2.996 8-2.996S18.5 3 11.5 3c-5.143 0-8 1.2-8 3.004Zm0 6c0 1.804 2.857 2.996 8 2.996c7 0 8-2.996 8-2.996S18.5 9 11.5 9c-5.143 0-8 1.2-8 3.004Zm0 6c0 1.804 2.857 2.996 8 2.996c7 0 8-2.996 8-2.996S18.5 15 11.5 15c-5.143 0-8 1.2-8 3.004ZM19.5 12s1-.975 1-3s-1-3-1-3m1-2c0 1.35-1 2-1 2m0 12s1-.975 1-3s-1-3-1-3m1 8c0-1.35-1-2-1-2"/>`),
		g.Group(children),
	)
}